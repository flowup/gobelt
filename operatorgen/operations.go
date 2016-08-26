package operatorgen

import (
	"os"
	"path/filepath"

	"github.com/flowup/gobelt"
	"github.com/flowup/gogen"
)

// TemplateData is a data structure for the operations template
type TemplateData struct {
	Package   string
	ModelName string
}

// TType template structures
type TType struct {
	code int
}

// Generate parses all given files by args and generates operations
// structures for structures annotated by @operations build tag
func Generate(args []string) error {
	return gobelt.Generate(args, func(build *gogen.Build, filePath, dir string) error {
		// retrieve the file from the build
		file := build.Files[filePath]

		data := TemplateData{
			Package: file.Package(),
		}

		for stName := range file.Structs().Filter("@operations") {
			_, err := os.Create(filepath.Join(dir, stName+".operations.gen.go"))
			if err != nil {
				return err
			}

			data.ModelName = stName
		}

		return nil
	})
}
