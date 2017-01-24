package observablegen

import (
  "os"
  "path/filepath"

  "github.com/flowup/gogen"
  "github.com/flowup/gobelt"
	"github.com/flowup/backend-services/filecache"
)

// TemplateData is a data structure for the observable template
type TemplateData struct {
  Package   string `template:"Package"`
  ModelName string `template:"TType"`
}

// template structures
type TType struct {}

// Generate parses all given files by args and generates observable
// structures for structures annotated by @observable build tag
func Generate(args []string) error {
  return gobelt.Generate(args, func(build *gogen.Build, filePath, dir string) error {
    // retrieve the file from the build
    file := build.File(filepath.Base(filePath))

    data := TemplateData{
      Package: file.Package(),
    }

		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

    // open template file
    template := filecache.Cache.LoadFile(gobelt.GetTemplatePath("observablegen") + "/template.go")

    for stName := range file.Structs().Filter("@observable") {
      data.ModelName = stName

      // create out file
      out, err := os.Create(filepath.Join(pwd, stName + ".observable.gen.go"))
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
