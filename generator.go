package gobelt

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/flowup/gogen"
)

type GenerateCallback func(build *gogen.Build, file string, dir string) error

func Generate(files []string, cb GenerateCallback) error {
	for _, file := range files {
		// get the dir
		path := filepath.Dir(file)

		// get the build
		build, err := gogen.ParseFile(file)
		if err != nil {
			return err
		}

		if err = cb(build, file, path); err != nil {
			return err
		}
	}

	return nil
}

// GetTemplatePath returns path to the templates of specified generator.
// It automatically replaces forward slashes to backward on windows
func GetTemplatePath(gen string) string {
	templatePath := path.Join(os.Getenv("GOPATH"), "/src/github.com/flowup/gobelt", gen)
	if runtime.GOOS == "windows" {
		templatePath = strings.Replace(templatePath, "/", "\\", -1)
	}

	return templatePath
}

// ExecuteTemplate writes a template into the given writer.
// Template data will be replaced by given tags
func ExecuteTemplate(templateBytes []byte, out io.Writer, data interface{}) error {
	st := reflect.ValueOf(data).Elem()

	templateData := string(templateBytes)

	for i := 0; i < st.NumField(); i++ {
		valueField := st.Field(i)       // valueField.Interface{}
		typeField := st.Type().Field(i) // typeField.Name
		templateTag := typeField.Tag.Get("template")

		value := valueField.Interface().(string)

		if templateTag == "Package" {
			// replace upper package by Package}

			// TODO: place this on top
			reg := regexp.MustCompile("package \\w+")

			templateData = string(reg.ReplaceAll([]byte(templateData), []byte("package " + value)))

			// continue to next iteration as this is special case
			continue
		}

		// replace given tag by the value
		templateData = strings.Replace(templateData, templateTag, value, -1)
	}

	_, err := out.Write([]byte(templateData))

	return err
}
