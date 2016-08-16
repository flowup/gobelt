package observablegen

import (
  "os"
  "path/filepath"

  "github.com/flowup/gogen"
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
  for _, arg := range args {
    // get the dir
    path := filepath.Dir(arg)

    // get the build
    build, err := gogen.ParseFile(arg)
    if err != nil {
      return err
    }

    // retrieve the file from the build
    file := build.Files[arg]

    data := TemplateData{
      Package: file.Package(),
    }

    for stName, _ := range file.Structs().Filter("@observable") {
      _, err := os.Create(filepath.Join(path, stName + ".observable.gen.go"))
      if err != nil {
        return err
      }

      data.ModelName = stName
    }
  }

  return nil
}
