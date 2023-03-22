package reader

import (
	"io"

	"gopkg.in/yaml.v3"
)

// NewYAML creates new yaml reader
func NewYAML(r io.Reader) *YAML {
	var rd YAML
	if r != nil {
		yaml.NewDecoder(r).Decode(&rd.d)
	}
	return &rd
}

// YAML reads config from yaml file
type YAML struct {
	d map[string]string
}

// Read reads a config
func (r *YAML) Read(name string) ([]byte, error) {
	p, ok := r.d[name]
	if !ok {
		return nil, errNotFound
	}
	return []byte(p), nil
}
