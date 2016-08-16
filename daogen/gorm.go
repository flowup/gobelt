package daogen

import (
  "os"
  "path/filepath"
  "runtime"
  "strings"
  "text/template"

  "github.com/azer/snakecase"
  "github.com/flowup/gogen"
)

var (
  headerTemplate = template.Must(template.New("header").Parse(`
package {{.Package}}

import (
	"github.com/jinzhu/gorm" {{.ProjectImport}}
)

`))

  serviceTemplate = template.Must(template.New("service").Parse(`

/*
Base DAO operations start
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
Base CRUD operations.
*/


// Create will create single {{.ModelName}} in database.
func (dao *{{.DAOName}}) Create(m *{{.ModelPackage}}{{.ModelName}}) {
  dao.db.Create(m)
}

// Read will find all DB records matching
// values in a model given by parameter
func (dao *{{.DAOName}}) Read(m *{{.ModelPackage}}{{.ModelName}}) []{{.ModelPackage}}{{.ModelName}} {
  retVal := []{{.ModelPackage}}{{.ModelName}}{}
  dao.db.Where(m).Find(&retVal)

  return retVal
}

// ReadByID will find {{.ModelName}} by ID given by parameter
func (dao *{{.DAOName}}) ReadByID(id uint64) *{{.ModelPackage}}{{.ModelName}}{
  m := &{{.ModelPackage}}{{.ModelName}}{}
  if dao.db.First(&m, id).RecordNotFound() {
    return nil
  }

  return m
}


// Update will update a record of {{.ModelName}} in DB
func (dao *{{.DAOName}}) Update(m *{{.ModelPackage}}{{.ModelName}}, id uint64) *{{.ModelPackage}}{{.ModelName}}{
	oldVal := dao.ReadByID(id)
	if oldVal == nil {
		return nil
	}

	dao.db.Model(&oldVal).Updates(m)
	return oldVal
}

// Delete will soft-delete a single {{.ModelName}}
func (dao *{{.DAOName}}) Delete(m *{{.ModelPackage}}{{.ModelName}}) {
  dao.db.Delete(m)
}
`))

  simpleFieldTemplate = template.Must(template.New("simpleField").Parse(`
/*
{{.FieldName}} related operations.
*/

// ReadBy{{.FieldName}} will find all records
// matching the value given by parameter
func (dao *{{.DAOName}}) ReadBy{{.FieldName}} (m {{.FieldType}}) []{{.ModelPackage}}{{.ModelName}} {
  retVal := []{{.ModelPackage}}{{.ModelName}}{}
  dao.db.Where(&{{.ModelPackage}}{{.ModelName}}{ {{.FieldName}} : m }).Find(&retVal)

  return retVal
}

// DeleteBy{{.FieldName}} deletes all records in database with
// {{.FieldName}} the same as parameter given
func (dao *{{.DAOName}}) DeleteBy{{.FieldName}} (m {{.FieldType}}) {
  dao.db.Where(&{{.ModelPackage}}{{.ModelName}}{ {{.FieldName}} : m }).Delete(&{{.ModelPackage}}{{.ModelName}}{})
}

// EditBy{{.FieldName}} will edit all records in database
// with the same {{.FieldName}} as parameter given
// using model given by parameter
func (dao *{{.DAOName}}) EditBy{{.FieldName}} (m {{.FieldType}}, newVals *{{.ModelPackage}}{{.ModelName}}) {
  dao.db.Table("{{.TableName}}").Where(&{{.ModelPackage}}{{.ModelName}}{ {{.FieldName}} : m }).Updates(newVals)
}

// Set{{.FieldName}} will set {{.FieldName}}
// to a value given by parameter
func (dao *{{.DAOName}}) Set{{.FieldName}} (m *{{.ModelPackage}}{{.ModelName}}, newVal {{.FieldType}}) *{{.ModelPackage}}{{.ModelName}} {
  m.{{.FieldName}} = newVal
  record := dao.ReadByID(uint64(m.ID))

  dao.db.Model(&record).Updates(m)

  return record
}
`))

  sliceFieldTemplate = template.Must(template.New("sliceField").Parse(`

/*
{{.FieldName}} related operations.
*/

func (dao *{{.DAOName}}) Add{{.FieldName}}Association (m *{{.ModelPackage}}{{.ModelName}}, asocVal {{.FieldType}}) *{{.ModelPackage}}{{.ModelName}} {
  dao.db.Model(&m).Association("{{.FieldName}}").Append(asocVal)

  return m
}

func (dao *{{.DAOName}}) Remove{{.FieldName}}Association (m *{{.ModelPackage}}{{.ModelName}}, asocVal {{.FieldType}}) *{{.ModelPackage}}{{.ModelName}} {
  dao.db.Model(&m).Association("{{.FieldName}}").Delete(asocVal)

  return m
}
`))

  footerTemplate = template.Must(template.New("footer").Parse(`

`))
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
  // for each passed file
  for _, arg := range args {
    // get only the file name
    name := strings.Split(filepath.Base(arg), ".")[0]
    // get the dir
    path := filepath.Dir(arg)
    absolutePath, _ := filepath.Abs(arg)
    pwd, err := os.Getwd()
    pwd = filepath.Base(pwd)
    out, err := os.Create(filepath.Join(path, name + ".dao.gen.go"))
    if err != nil {
      return err
    }

    // get the build
    build, err := gogen.ParseFile(arg)
    if err != nil {
      return err
    }
    importString := strings.TrimLeft(absolutePath, os.Getenv("GOPATH")+"src")
    if runtime.GOOS == "windows" {
      importString = strings.Replace(importString, "\\", "/", -1)
    }

    importString = strings.TrimRight(importString, "/" + name + ".go")

    // retrieve the file from the build
    file := build.Files[arg]

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

    // add header to the test file
    headerTemplate.Execute(out, data)

    // iterate over structures
    for stName, stVal := range file.Structs() {
      // update suite name
      data.ServiceName = stName
      data.ModelName = stName
      data.DAOName = data.ModelName + "DAO"
      data.TableName = snakecase.SnakeCase(stName) + "s"
      // add the suite
      serviceTemplate.Execute(out, data)

      for _, fieldVal := range stVal.Fields() {

        var typeType int
        data.FieldName = fieldVal.Name()
        data.FieldType, typeType = fieldVal.Type()
        //if it is not a gorm ID or one of
        // the time parameters execute field template
        if data.FieldName != "ID" &&
          data.FieldName != "CreatedAt" &&
          data.FieldName != "UpdatedAt" &&
          data.FieldName != "DeletedAt" {
          switch typeType {
          case gogen.PrimitiveType:
            simpleFieldTemplate.Execute(out, data)
          case gogen.SliceType:
          //sliceFieldTemplate.Execute(out, data) // support for slices is not yet finished
          }
        }
      }

      // add suite test execution
      footerTemplate.Execute(out, data)
    }
  }

  return nil
}
