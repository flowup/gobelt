package daogen

import (
	"text/template"
	"strings"
	"path/filepath"
	"os"
	"github.com/flowup/gogen"
)

var (
	headerTemplate = template.Must(template.New("header").Parse(`

`))

	serviceTemplate = template.Must(template.New("service").Parse(`

`))

	footerTemplate = template.Must(template.New("footer").Parse(`

`))
)

type TemplateData struct {
	Package string
	ServiceName string
}

func GenerateGorm(args []string) error {
	// for each passed file
	for _, arg := range args {
		// get only the file name
		name := strings.Split(filepath.Base(arg), ".")[0]
		// get the dir
		path := filepath.Dir(arg)
		out, err := os.Create(filepath.Join(path, name + ".service.go"))
		if err != nil {
			return err
		}

		// get the build
		build, err := gogen.ParseFile(arg)
		if err != nil {
			return err
		}

		// retrieve the file from the build
		file := build.Files[arg]

		// initialize the data structure
		data := TemplateData{
			Package: file.Package(),
			ServiceName: "",
		}

		// add header to the test file
		headerTemplate.Execute(out, data)

		// iterate over structures
		for stName, _ := range file.Structs() {
			// update suite name
			data.ServiceName = stName
			// add the suite
			serviceTemplate.Execute()

			// add suite test execution
			footerTemplate.Execute(out, data)
		}
	}

	return nil
}