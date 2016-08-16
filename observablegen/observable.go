package observablegen

import (
  "os"
  "path/filepath"

  "github.com/flowup/gogen"
  "github.com/flowup/gobelt"
)

// TemplateData is a data structure for the observable template
type TemplateData struct {
  Package   string
  ModelName string
}

// template structures
type __T__ struct {}

// Generate parses all given files by args and generates observable
// structures for structures annotated by @observable build tag
func Generate(args []string) error {
  return gobelt.Generate(args, func(build *gogen.Build, filePath, dir string) error {
    // retrieve the file from the build
    file := build.Files[filePath]

    data := TemplateData{
      Package: file.Package(),
    }

    for stName := range file.Structs().Filter("@observable") {
      _, err := os.Create(filepath.Join(dir, stName + ".observable.gen.go"))
      if err != nil {
        return err
      }

      data.ModelName = stName
    }

    return nil
  })
}
