package daogen

import (
	"text/template"
	"strings"
	"path/filepath"
	"os"
	"github.com/flowup/gogen"
  "runtime"
)

var (
	headerTemplate = template.Must(template.New("header").Parse(`
package dao

import (
	"github.com/jinzhu/gorm"
	"{{.ProjectImport}}/{{.Package}}"
)

`))

	serviceTemplate = template.Must(template.New("service").Parse(`

/*
@Init
*/

// {{.DAOName}} is a data access object to a database containing {{.ModelName}}s
type {{.DAOName}} struct {
	db *gorm.DB
}

// New{{.DAOName}} creates a new Data Access Object for the
// {{.ModelName}} model.
func New{{.DAOName}} (db *gorm.DB) *{{.DAOName}} {
	return &{{.DAOName}}{
		db: db,
	}
}

/*
@CRUD
*/


// Create will create single {{.ModelName}} in database.
func (dao *{{.DAOName}}) Create(m *{{.Package}}.{{.ModelName}}) {
  dao.db.Create(m)
}

func (dao *{{.DAOName}}) Read(m *{{.Package}}.{{.ModelName}}) {
}

// FindByID will find {{.ModelName}} by ID given by parameter
func (dao *{{.DAOName}}) FindByID(id uint64) *{{.Package}}.{{.ModelName}}{
  var m *{{.Package}}.{{.ModelName}}
  if dao.db.First(&m, id).RecordNotFound() {
    return nil
  }

  return m
}


// Update will update a record of {{.ModelName}} in DB
func (dao *{{.DAOName}}) Update(m *{{.Package}}.{{.ModelName}}) {
}

// Delete will soft-delete a single {{.ModelName}}
func (dao *{{.DAOName}}) Delete(m *{{.Package}}.{{.ModelName}}) {
  dao.db.Delete(m)
}
`))

	footerTemplate = template.Must(template.New("footer").Parse(`

`))
)

type TemplateData struct {
	Package string
	ServiceName string
	ModelName string
  ProjectImport string
  DAOName string
}

func GenerateGorm(args []string) error {
	// for each passed file
	for _, arg := range args {
		// get only the file name
		name := strings.Split(filepath.Base(arg), ".")[0]
		// get the dir
		path := filepath.Dir(arg)
    absolutePath, _ := filepath.Abs(arg)
		out, err := os.Create(filepath.Join(path, name + ".service.go"))
		if err != nil {
			return err
		}

		// get the build
		build, err := gogen.ParseFile(arg)
		if err != nil {
			return err
		}

    var splitPaths []string
    var importString string
    if runtime.GOOS == "windows" {
      splitPaths = strings.SplitAfter(absolutePath, "src\\")
      importString = splitPaths[len(splitPaths) - 1]
      importString = strings.Replace(importString, "\\", "/", -1)
    } else {
      splitPaths = strings.SplitAfter(absolutePath, "src/")
      importString = splitPaths[len(splitPaths) - 1]
    }

    importString = strings.TrimRight(importString, "/"+name+".go")

		// retrieve the file from the build
		file := build.Files[arg]

		// initialize the data structure
		data := TemplateData{
			Package: file.Package(),
			ServiceName: "",
      ProjectImport: importString,
      // currently ProjectImport is parsed from path,
      // should be parsed using gogen
      // (could not work if a project has src/ directory in it)
		}

		// add header to the test file
		headerTemplate.Execute(out, data)

		// iterate over structures
		for stName, _ := range file.Structs() {
			// update suite name
			data.ServiceName = stName
      data.ModelName = stName
      data.DAOName = data.ModelName + "DAO"
			// add the suite
			serviceTemplate.Execute(out, data)

			// add suite test execution
			footerTemplate.Execute(out, data)
		}
	}

	return nil
}