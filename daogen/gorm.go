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
package {{.Package}}

import (
	"github.com/jinzhu/gorm"
	"errors"
	"{{.ModelImport}}"
)

/*
@Init
*/

type {{.ModelName}}DAO struct {
	db *gorm.DB
}

// New{{.Model}}DAO creates a new Data Access Object for the
// {{.ModelName}} model.
func New{{.ModelName}}DAO (db *gorm.DB) *{{.ModelName}}Service {
	return &{{.ModelName}}DAO{
		db: db,
	}
}
`))

	serviceTemplate = template.Must(template.New("service").Parse(`
/*
@CRUD
*/

func (dao *{{.ModelName}}DAO) Create(m *{{.ModelImport}}) ({
}

func (dao *{{.ModelName}}DAO) Read(m *{{.ModelImport}}) ({
}

func (dao *{{.ModelName}}DAO) Update(m *{{.ModelImport}}) ({
}

func (dao *{{.ModelName}}DAO) Delete(m *{{.ModelImport}}) ({
}
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
			serviceTemplate.Execute(out, data)

			// add suite test execution
			footerTemplate.Execute(out, data)
		}
	}

	return nil
}