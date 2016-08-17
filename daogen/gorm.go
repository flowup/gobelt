package daogen

import (
  "os"
  "path/filepath"
  "runtime"
  "strings"
  "github.com/azer/snakecase"
  "github.com/flowup/gogen"
  "io/ioutil"
  "github.com/flowup/gobelt"
)

type TemplateData struct {
  Package       string
  ModelPackage  string
  ServiceName   string
  ModelName     string
  ProjectImport string
  DAOName       string
  TableName     string
  FieldName     string
  FieldType     string
}

func GenerateGorm(args []string) error {
  return gobelt.Generate(args, func(build *gogen.Build, filePath, dir string) error {
    name := strings.Split(filepath.Base(filePath), ".")[0]
    absolutePath, _ := filepath.Abs(dir)

    // create file for output
    out, err := os.Create(filepath.Join(dir, name + ".dao.gen.go"))
    if err != nil {
      return err
    }
    defer out.Close()


    pwd, err := os.Getwd()
    pwd = filepath.Base(pwd)

    // parse import string using $GOPATH
    importString := strings.TrimLeft(absolutePath, os.Getenv("GOPATH")+"src")
    if runtime.GOOS == "windows" {
      importString = strings.Replace(importString, "\\", "/", -1)
    }

    importString = strings.TrimRight(importString, "/" + name + ".go")

    // retrieve the file from the build
    file := build.Files[filePath]

    var modelPackage string
    if pwd == file.Package() {
      modelPackage = ""
      importString = ""
    } else {
      modelPackage = file.Package() + "."
      importString = "\n  \"" + importString + "\""
    }

    // initialize the data structure
    data := TemplateData{
      Package:       pwd,
      ModelPackage:  modelPackage,
      ServiceName:   "",
      ProjectImport: importString,
    }

    // compose template files path and open them
    openPath := os.Getenv("GOPATH") + "/src/github.com/flowup/gobelt/daogen/"
    if runtime.GOOS == "windows" {
      openPath = strings.Replace(openPath, "/", "\\", -1)
    }

    templateBase, err := os.Open(openPath + "templateBase.go")
    if err != nil {
      return err
    }
    defer templateBase.Close()


    var baseString string
    baseBytes, err := ioutil.ReadAll(templateBase)
    if err != nil {
      return err
    }

    // skip the import and package section
    lines := strings.Split((string)(baseBytes), "\n)")
    for i := 1; i < len(lines); i++ {
      baseString += lines[i]
    }

    templatePrimitive, err := os.Open(openPath + "templatePrimitive.go")
    if err != nil {
      return err
    }
    defer templatePrimitive.Close()

    primitiveRead, err := ioutil.ReadAll(templatePrimitive)
    if err != nil {
      return err
    }
    primitiveString := strings.TrimLeft((string)(primitiveRead), "package daogen\n")


    /*
    templateSlice, err := os.Open(openPath + "templateSlice.go")
    if err != nil {
      return err
    }
    defer templateSlice.Close()
    */

    // write the header containing package and imports into output file
    out.Write(([]byte)("package " + data.Package + "\n\nimport(\n  \"github.com/jinzhu/gorm\""   + data.ProjectImport + "\n)"))

    var outputString string
    // iterate over structures
    for stName, stVal := range file.Structs() {
      // reset outputString to base template
      outputString = baseString

      // update suite name
      data.ServiceName = stName
      data.ModelName = stName
      data.DAOName = data.ModelName + "DAO"
      data.TableName = snakecase.SnakeCase(stName) + "s"

      for _, fieldVal := range stVal.Fields() {

        var typeType int
        data.FieldName = fieldVal.Name()
        data.FieldType, typeType = fieldVal.Type()
        // if it is not a gorm ID or one of
        // the time parameters execute field template
        if data.FieldName != "ID" &&
          data.FieldName != "CreatedAt" &&
          data.FieldName != "UpdatedAt" &&
          data.FieldName != "DeletedAt" {

          var fieldOps string

          switch typeType {
          case gogen.PrimitiveType:
            // compose functions for primitive types
            fieldOps = strings.Replace(primitiveString, "__PrimitiveType__", data.FieldType, -1)
            fieldOps = strings.Replace(fieldOps, "__FieldPrimitive__", data.FieldName, -1)
          case gogen.SliceType:
            // support for slices is not yet finished
          }
          outputString += (fieldOps)
        }
      }
      // replace template names with the names of current structure
      outputString = strings.Replace(outputString, "__DAOName__", data.DAOName, -1)
      outputString = strings.Replace(outputString , "__ReferenceModel__", data.ModelPackage + data.ModelName, -1)
      outputString = strings.Replace(outputString, "__P__", data.Package, -1)

      // write code for parsed struct to output file
      out.Write(([]byte)(outputString))
    }

    return nil
  })
}
