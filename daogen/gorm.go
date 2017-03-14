package daogen

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/azer/snakecase"
	"github.com/flowup/backend-services/filecache"
	"github.com/flowup/gobelt"
	"github.com/flowup/gogen"
	"fmt"
	"regexp"
)

// TemplateData is used to hold data about given models
type TemplateData struct {
	Package       string
	ModelPackage  string
	ModelName     string
	ProjectImport string
	DAOName       string
	TableName     string
	FieldName     string
	FieldType     string
}

// HeaderSeparator separates headers from methods in templates
const(
	HeaderSeparator = "/* END OF HEADER */"
	ImplementationPrefix = `\(dao [^\)]*\) `
)


// GenerateGorm will return a callback function that will generate a gorm based
// data access object for given models
func GenerateGorm(args []string) error {
	return gobelt.Generate(args, func(build *gogen.Build, filePath, dir string) error {
		name := strings.Split(filepath.Base(filePath), ".")[0]
		absolutePath, _ := filepath.Abs(dir)

		// create file for output
		pwdBase, err := os.Getwd()
		pwd := filepath.Base(pwdBase)
		out, err := os.Create(filepath.Join(pwdBase, name+".dao.gen.go"))
		if err != nil {
			return err
		}
		defer out.Close()

		// parse import string using $GOPATH
		//importString := strings.TrimLeft(absolutePath, os.Getenv("GOPATH")+"src")
		importString := absolutePath[len(os.Getenv("GOPATH")+"/src/"):]
		if runtime.GOOS == "windows" {
			importString = strings.Replace(importString, "\\", "/", -1)
		}
		// retrieve the file from the build
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

		// initialize the data structure
		data := TemplateData{
			Package:       pack,
			ModelPackage:  modelPackage,
			ProjectImport: importString,
		}

		// compose template files path and open them


		matcher := regexp.MustCompile(ImplementationPrefix + `[A-Z]\w*\([^\)]*\) \([^\)]*\)`)
		var baseString string

		templatePath := gobelt.GetTemplatePath("daogen")
		baseBytes := filecache.Cache.LoadFile(templatePath + "/template_base.go")

		// skip the import and package section
		lines := strings.Split((string)(baseBytes), "\n)")
		for i := 1; i < len(lines); i++ {
			baseString += lines[i]
		}

		readByID := filecache.Cache.LoadFile(templatePath + "/template_readByID.go")
		readByIDString := strings.Split(string(readByID), HeaderSeparator)[1]

		readByIDEmbedded := filecache.Cache.LoadFile(templatePath + "/template_readByIDEmbedded.go")
		readByIDEmbeddedString := strings.Split(string(readByIDEmbedded), HeaderSeparator)[1]

		primitiveRead := filecache.Cache.LoadFile(templatePath + "/template_primitive.go")
		primitiveString := strings.Split(string(primitiveRead), HeaderSeparator)[1]
		//strings.TrimLeft((string)(primitiveRead), "package daogen\n")

		uintIDMockRead := filecache.Cache.LoadFile(templatePath + "/template_uintIDMock.go")
		uintIDMockString := strings.Split(string(uintIDMockRead), HeaderSeparator)[1]

		sliceRead := filecache.Cache.LoadFile(templatePath + "/template_slice.go")
		sliceString := strings.Split(string(sliceRead), HeaderSeparator)[1]
		//strings.TrimLeft((string)(sliceRead), "package daogen\n")

		stringIDMockRead := filecache.Cache.LoadFile(templatePath + "/template_stringIDMock.go")
		stringIDMockString := strings.Split(string(stringIDMockRead), HeaderSeparator)[1]

		structRead := filecache.Cache.LoadFile(templatePath + "/template_struct.go")
		structString := strings.Split(string(structRead), HeaderSeparator)[1]
		//strings.TrimLeft((string)(structRead), "package daogen\n")

		var neededPackages []string
		var outputStrings []string
		// iterate over structures
		for stName, stVal := range file.Structs().Filter("@dao") {
			// reset outputString to base template
			outputString := baseString
			idTypeString := "uint"
			embStructWithIDName := ""

			// update suite name
			data.ModelName = stName
			data.DAOName = data.ModelName + "DAO"
			data.TableName = snakecase.SnakeCase(stName) + "s"

			for fieldName, fieldVal := range stVal.Fields() {
				if len(fieldName) != 0 && fieldName[0] == '_' {
					embStr := file.Struct(fieldName[1:])
					if embStr != nil {
						for embFiName, embFi := range embStr.Fields() {
							if embFiName == "ID" {
								embStructWithIDName = embStr.Name()
								idTypeString, _ = embFi.Type()
								if idTypeString == "string" {
									neededPackages = append(neededPackages, `"strconv"`)
								}
								outputString += readByIDEmbeddedString
								break
							}
						}
					}
				} else {
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
						case gogen.SelectorType:
							neededPackages = append(neededPackages, strings.Split(data.FieldType, ".")[0])
							fallthrough
						case gogen.StructType:
							if data.FieldName != "" {
								// compose functions for struct types
								fieldOps = strings.Replace(structString, "FieldStruct", data.FieldName, -1)
								fieldOps = strings.Replace(fieldOps, "AuxModel", data.FieldType, -1)
							}
						case gogen.PrimitiveType:
							if data.FieldName != "" {
								// compose functions for primitive types
								fieldOps = strings.Replace(primitiveString, "PrimitiveType", data.FieldType, -1)
								fieldOps = strings.Replace(fieldOps, "FieldPrimitive", data.FieldName, -1)
							}
						case gogen.SliceType:
							if data.FieldName != "" {
								// compose functions for array types
								fieldOps = strings.Replace(sliceString, "AuxModel", data.ModelPackage+data.FieldType, -1)
								fieldOps = strings.Replace(fieldOps, "FieldSlice", data.FieldName, -1)
							}
						}
						outputString += fieldOps
					}
				}
			}
			if embStructWithIDName == "" {
				outputString += readByIDString
			}

			if idTypeString == "uint" {
				outputString += uintIDMockString
			} else if idTypeString == "string" {
				outputString += stringIDMockString
			}
			// replace template names with the names of current structure

			outputString = strings.Replace(outputString, "reference_models", data.TableName, -1)
			outputString = strings.Replace(outputString, "ReferenceModelIDType", idTypeString, -1)
			outputString = strings.Replace(outputString, "DAONameString", data.DAOName, -1)
			outputString = strings.Replace(outputString, "DAONameEmbedded", data.DAOName, -1)
			outputString = strings.Replace(outputString, "DAOName", data.DAOName, -1)
			outputString = strings.Replace(outputString, "AuxModelEmbedded", embStructWithIDName, -1)
			outputString = strings.Replace(outputString, "ReferenceModelStringID", data.ModelPackage+data.ModelName, -1)
			outputString = strings.Replace(outputString, "ReferenceModelEmbedded", data.ModelPackage+data.ModelName, -1)
			outputString = strings.Replace(outputString, "ReferenceModel", data.ModelPackage+data.ModelName, -1)

			decls := matcher.FindAllString(outputString, -1)
			fmt.Println(decls)
			if decls != nil {
				outputString += "\ntype " + data.DAOName + "Interface interface {\n"
				cut := len("(dao *" + data.DAOName + ")")
				for _, dec := range decls {
					fmt.Println(dec[cut:])
					outputString += "\t" + dec[len(ImplementationPrefix) + 2:] + "\n"
				}
				outputString += "}"
			}

			outputStrings = append(outputStrings, outputString)
		}
		generatedStr := "package " + data.Package + "\n\nimport(\n  \"github.com/jinzhu/gorm\"\n  \"time\"" + data.ProjectImport + "\n  "
		for _, pack := range neededPackages {
			if i := file.Import(pack); i != nil && pack != "time" {
				generatedStr += "  " + i.String() + "\n"
			}
			if pack == `"strconv"` {
				generatedStr += "  " + pack + "\n"
			}
		}
		generatedStr += ")\n"
		for _, str := range outputStrings {
			generatedStr += str
		}

		// write code for parsed struct to output file
		out.Write(([]byte)(generatedStr))
		return nil
	})
}
