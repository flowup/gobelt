package suitegen

import (
  "os"
  "path/filepath"
  "strings"
  "text/template"

  "github.com/flowup/gogen"
)

var (
  // template with the package name and imports
  headerTemplate = template.Must(template.New("header").Parse(`
package {{.Package}}

import (
  "github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
)
`))

  // template with the suite structure
  suiteTemplate = template.Must(template.New("suite").Parse(`
type {{.SuiteName}}Suite struct {
	suite.Suite
}

func (s *{{.SuiteName}}Suite) SetupSuite() {

}

func (s *{{.SuiteName}}Suite) TearDownSuite() {

}

func (s *{{.SuiteName}}Suite) SetupTest() {

}

func (s *{{.SuiteName}}Suite) TearDownTest() {

}
`))

  // Template of each function
  functionTemplate = template.Must(template.New("function").Parse(`
func (s *{{.SuiteName}}Suite) Test{{.FuncName}}() {
  assert.Fail(s.T(), "Test 'Test{{.FuncName}}' is not implemented")
}
`))

  // template with suite start and options
  footerTemplate = template.Must(template.New("footer").Parse(`
func Test{{.SuiteName}}Suite(t *testing.T) {
	suite.Run(t, &{{.SuiteName}}Suite{})
}
`))
)

//
type TemplateData struct {
  Package   string // name of the package
  SuiteName string // suite name (without Suite)
  FuncName  string // currently iterated function
}

// Generate accepts a list of files or directories
// that is being used to create test suites.
func Generate(args []string) error {
  // TODO: accept folders
  for _, arg := range args {
    // get only the file name
    name := strings.Split(filepath.Base(arg), ".")[0]
    // get the dir
    path := filepath.Dir(arg)
    out, err := os.Create(filepath.Join(path, name + "_test.go"))
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
      Package:   file.Package(),
      SuiteName: "",
      FuncName:  "",
    }

    // add header to the test file
    headerTemplate.Execute(out, data)

    // iterate over structures
    for stName, st := range file.Structs() {
      // update suite name
      data.SuiteName = stName
      // add the suite
      suiteTemplate.Execute(out, data)

      // iterate over their functions
      for methodName, _ := range st.Methods() {
        // update name of the method
        data.FuncName = methodName

        functionTemplate.Execute(out, data)
      }

      // add suite test execution
      footerTemplate.Execute(out, data)
    }
  }

  return nil
}
