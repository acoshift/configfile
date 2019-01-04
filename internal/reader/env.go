package reader

import (
	"os"
	"strings"
)

// NewEnv creates new env reader
func NewEnv() *Env {
	var r Env
	return &r
}

// Env reads config from env
type Env struct{}

func (r *Env) Read(name string) ([]byte, error) {
	name = strings.ToUpper(name)
	p, ok := os.LookupEnv(name)
	if !ok {
		return nil, errNotFound
	}
	return []byte(p), nil
}
