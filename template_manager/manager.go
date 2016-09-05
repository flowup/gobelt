package templateManager

import (
	"runtime"
	"strings"
	"os"
	"io/ioutil"
)

// instance of a manager
var instance *Manager

// Manager is a structure holding cached templates
type Manager struct {
	loadedTemplates map[string][]byte
}

// GetInstance will get an instance of manager.
// If it does not yet exist, it will create it.
func GetInstance() *Manager {
	if instance == nil {
		instance = &Manager{
			loadedTemplates: make(map[string][]byte),
		}
	}
	return instance
}

// LoadTemplateNoCache will load a template and NOT save it in
// cache.
func (m *Manager) LoadTemplateNoCache(path string) []byte {
	filePath := os.Getenv("GOPATH") + "/src/github.com/flowup/gobelt/" + path
	if runtime.GOOS == "windows" {
		filePath = strings.Replace(filePath, "/", "\\", -1)
	}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return bytes
}

// LoadTemplate will look into cache if the requested
// template is not loaded already, if it is, it will return it
// without any further ado, if it is not, it will load it,
// save it to cache and return it.
func (m *Manager) LoadTemplate(path string) []byte{
	if bytes, ok := m.loadedTemplates[path]; ok == true {
		return bytes
	}
	m.loadedTemplates[path] = m.LoadTemplateNoCache(path)
	return m.loadedTemplates[path]
}
