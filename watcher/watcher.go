package watcher

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"flowdock.eu/flowup/services/filecache"

	"github.com/flowup/gobelt"
	"github.com/flowup/gogen"
	"github.com/fsnotify/fsnotify"
)

// TemplateData is a data structure
type TemplateData struct {
	Package   string `template:"Package"`
	ModelName string `template:"TType"`
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// LoadAllTemplates is loading all files
func LoadAllTemplates(templates []string) map[string][]byte {
	templatesMap := make(map[string][]byte)
	for _, template := range templates {
		// Cutting off ".generic.go" and path for saving as key, so key is the name of the template, in cache the is former path and full name
		splittedName := strings.Split(template, ".")
		templateName := (splittedName[0])
		templatesMap[path.Base(templateName)] = filecache.Cache.LoadFile(template)
	}
	return templatesMap

}

// WalkSubdirectories is traversing the directories
func WalkSubdirectories(directory string) ([]string, []string) {

	fileList := []string{}
	directoryList := []string{}
	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			directoryList = append(directoryList, path)
		} else if !f.IsDir() && strings.HasSuffix(f.Name(), ".generic.go") {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return directoryList, fileList
}

// ParseFile is parsing file
func ParseFile(pathToFile string) []string {

	r, _ := regexp.Compile(`//\s*@template\s+[a-zA-Z]*`)

	dat, err := ioutil.ReadFile(pathToFile)
	panicIf(err)
	generatorTags := (removeCommentTag(removeDuplicates(r.FindAllString(string(dat), -1))))
	return generatorTags

}

func removeDuplicates(list []string) []string {
	mapDuplicates := make(map[string]bool, len(list))
	for _, x := range list {
		mapDuplicates[x] = true
	}
	result := make([]string, 0, len(mapDuplicates))
	for x := range mapDuplicates {
		result = append(result, x)
	}
	return result
}

func removeCommentTag(list []string) []string {
	result := make([]string, 0, len(list))
	r, _ := regexp.Compile(`//\s*`)
	for _, x := range list {
		decorator := r.ReplaceAllString(x, "")
		result = append(result, decorator)
	}
	return result
}

// Generate parses all given files by args and generates observable
// structures for structures annotated by @observable build tag
func Generate(file []string, fileName string, templateMap map[string][]byte) error {
	return gobelt.Generate(file, func(build *gogen.Build, filePath, dir string) error {
		// retrieve the file from the build

		file := build.File(path.Base(filePath))

		data := TemplateData{
			Package: file.Package(),
		}

		for stName, st := range file.Structs().Filter("@template") {
			data.ModelName = stName
			var generic string
			for _, value := range st.Tags().GetAll() {
				if value.Has("name") {
					generic, _ = value.Get("name")
				}
			}

			template, ok := templateMap[generic]
			if !ok {
				fmt.Println(generic + " template was not found!")
				continue
			}
			// create out file
			out, err := os.Create(filepath.Join(dir, fileName+"."+stName+"."+generic+".gen.go"))
			if err != nil {
				return err
			}

			err = gobelt.ExecuteTemplate(template, out, &data)
			if err != nil {
				return err
			}
			out.Close()

		}
		return nil
	})
}

// Watch is main process in watcher and it is processing watching loop
func Watch() {

	rootDirectory, err := os.Getwd()
	panicIf(err)

	directories, templates := WalkSubdirectories(rootDirectory)
	templateMap := LoadAllTemplates(templates)

	watcher, err := fsnotify.NewWatcher()
	panicIf(err)

	for _, directory := range directories {
		watcher.Add(directory)
	}

	for {
		select {
		case ev := <-watcher.Events:
			if !strings.HasSuffix(ev.Name, ".go") {
				// these are temporary files or others created by IDE's
				continue

			} else if strings.HasSuffix(ev.Name, "generic.go") {
				//fmt.Println("Triggered GENERIC change ", ev.Name)

				splittedName := strings.Split(ev.Name, ".")
				templateName := (splittedName[0])
				templateMap[path.Base(templateName)] = filecache.Cache.UpdateFile(ev.Name)

			} else if strings.HasSuffix(ev.Name, ".go") && !strings.HasSuffix(ev.Name, "generic.go") && !strings.HasSuffix(ev.Name, ".gen.go") {
				//fmt.Println("Searching for decorators @template TemplateName in .GO files", ev.Name)

				file := []string{ev.Name}

				splittedName := strings.Split(ev.Name, ".")
				filePath := (splittedName[0])
				fileName := path.Base(filePath)
				Generate(file, fileName, templateMap)

			}
		}
	}
}
