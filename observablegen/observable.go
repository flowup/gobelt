package observablegen

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/flowup/gogen"
)

var (
	observableTemplate = template.Must(template.New("observable").Parse(`
package {{.Package}}

type {{.ModelName}}SubscriberFunc func(val []*{{.ModelName}})

// {{.ModelName}}Observable is an observable for the {{.ModelName}}Subject
type {{.ModelName}}Observable struct {
	subscribers []*{{.ModelName}}Subscriber
}

// Subscribe will add the subscriber into the observable
func (o *{{.ModelName}}Observable) Subscribe(callback {{.ModelName}}SubscriberFunc) {
	o.subscribers = append(o.subscribers, &{{.ModelName}}Subscriber{callback})
}

// {{.ModelName}}Subscriber is an object that subscribes to the
// changes of {{.ModelName}}Subject.
type {{.ModelName}}Subscriber struct {
	callback {{.ModelName}}SubscriberFunc
}

// Next pushes the value to the subscriber
func (s *{{.ModelName}}Subscriber) Next(val []*{{.ModelName}}) {
	s.callback(val)
}

// {{.ModelName}}Subject is an entry point for changes of {{.ModelName}}
type {{.ModelName}}Subject struct {
	observable *{{.ModelName}}Observable
}

func New{{.ModelName}}Subject() *{{.ModelName}}Subject {
	return &{{.ModelName}}Subject{&{{.ModelName}}Observable{nil}}
}

// AsObservable returns the observable object of the {{.ModelName}}Subject
func (s *{{.ModelName}}Subject) AsObservable() *{{.ModelName}}Observable {
	return s.observable
}

// Next will push the given value to all subscribers of the
// underlying observable
func (s *{{.ModelName}}Subject) Next(val []*{{.ModelName}}) {
	for _, sub := range s.observable.subscribers {
		sub.Next(val)
	}
}
`))
)

// TemplateData is a data structure for the observable template
type TemplateData struct {
	Package   string
	ModelName string
}

func Generate(args []string) error {
	for _, arg := range args {
		// get the dir
		path := filepath.Dir(arg)

		// get the build
		build, err := gogen.ParseFile(arg)
		if err != nil {
			return err
		}

		// retrieve the file from the build
		file := build.Files[arg]

		data := TemplateData{
			Package: file.Package(),
		}

		for stName, _ := range file.Structs().Filter("@observable") {
			out, err := os.Create(filepath.Join(path, stName+".observable.gen.go"))
			if err != nil {
				return err
			}

			data.ModelName = stName

			observableTemplate.Execute(out, data)
		}
	}

	return nil
}
