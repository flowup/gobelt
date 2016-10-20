package restgen

import (
	"github.com/flowup/gobelt"
	"github.com/flowup/gogen"
)

func Generate(args []string) error {
	return gobelt.Generate(args, func(build *gogen.Build, filePath, dir string) error {

		return nil
	})
}
