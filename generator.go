package gobelt

import (
  "github.com/flowup/gogen"
  "path/filepath"
)

type GenerateCallback func(build *gogen.Build, file string, dir string) error

func Generate(files []string, cb GenerateCallback) error {
  for _, file := range files {
    // get the dir
    path := filepath.Dir(file)

    // get the build
    build, err := gogen.ParseFile(file)
    if err != nil {
      return err
    }

    if err = cb(build, file, path); err != nil {
      return err
    }
  }

  return nil
}