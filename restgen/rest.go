package restgen

import (
	"github.com/flowup/gobelt"
	"github.com/flowup/backend-services/filecache"
	"github.com/flowup/gogen"
	"path/filepath"
	"os"
	"strings"
)

func Generate(args []string) error {
	return gobelt.Generate(args, func(build *gogen.Build, filePath, dir string) error {
		name := strings.Split(filepath.Base(filePath), ".")[0]
		absolutePath, _ := filepath.Abs(dir)
		pwdBase, err := os.Getwd()
		pwd := filepath.Base(pwdBase)
		if err != nil {
			return err
		}
		out, err := os.Create(filepath.Join(pwdBase, name + ".rest.gen.go"))
		if err != nil {
			return err
		}
		defer out.Close()
		importString := absolutePath[len(os.Getenv("GOPATH") + "/src/"):]

		file := build.File(filepath.Base(filePath))
		var modelPackage string
		var pack string
		if absolutePath == pwdBase {
			pack = file.Package()
			modelPackage = ""
			importString = ""
		} else {
			pack = pwd
			modelPackage = file.Package() + "."
			importString = "\n  \"" + importString + "\""
		}

		templatePath := gobelt.GetTemplatePath("restgen")
		baseRead := filecache.Cache.LoadFile(templatePath + "/template_base.go")
		baseString := ""
		// skip the import and package section
		lines := strings.Split((string)(baseRead), "\n)")
		for i := 1; i < len(lines); i++ {
			baseString += lines[i]
		}


		var outputStrings []string

		for strName, str := range file.Structs().Filter("@rest") {
			outputString := baseString
			fieldOps := ""
			for _, field := range str.Fields() {
				fieldType, _ := field.Type()
				fieldName := field.Name()
				if fieldName != "ID" && fieldName != "CreatedAt" && fieldName != "UpdatedAt" && fieldName != "DeletedAt" {
					switch fieldType {
					case "string":
						fallthrough
					case "int64":
						fallthrough
					case "int":
						fieldOps += fieldName + " : load" + fieldType +"QueryParam(ctx, \"" + fieldName + "\"),\n\t\t\t"
					}
				}


			}
			outputString = strings.Replace(outputString, "ReferenceModel", modelPackage + strName, -1)
			outputString = strings.Replace(outputString, "//GENERATE_FIELDS", fieldOps, -1)

			outputStrings = append(outputStrings, outputString)
		}

		generatedString := "package " + pack + "\n\nimport(\n  \"github.com/kataras/iris\"" + importString + "\n)\n "
		for _, str := range outputStrings {
			generatedString += str
		}

		out.Write([]byte(generatedString))
		return nil
	})
}
