package operatorgen

import (
	"os"
	"path/filepath"

	"github.com/flowup/gobelt"
	"github.com/flowup/gogen"
)

// TemplateData is a data structure for the operations template
type TemplateData struct {
	Package   string `template:"Package"`
	ModelName string `template:"TType"`
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
		return FromFile(build.File(filePath), dir)
	})
}

func FromFile(file *gogen.File, targetDir string) error {
	data := TemplateData{
		Package: file.Package(),
	}

	template, err := os.Open(gobelt.GetTemplatePath("operatorgen/template.go"))
	if err != nil {
		return err
	}

	for stName := range file.Structs().Filter("@operations") {
		out, err := os.Create(filepath.Join(targetDir, stName+".operations.gen.go"))
		if err != nil {
			return err
		}

		data.ModelName = stName

		err = gobelt.ExecuteTemplate(template, out, &data)
		if err != nil {
			return err
		}

		out.Close()
	}

	return nil
}
