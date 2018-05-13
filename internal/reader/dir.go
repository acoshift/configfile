package reader

import (
	"io/ioutil"
	"path/filepath"
)

// Dir reads config from directory
type Dir struct {
	Base string
}

// Read reads a config
func (r *Dir) Read(name string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(r.Base, name))
}
