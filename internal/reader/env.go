package reader

import "os"

// NewEnv creates new env reader
func NewEnv() *Env {
	var r Env
	return &r
}

// Env reads config from env
type Env struct{}

func (r *Env) Read(name string) ([]byte, error) {
	p := os.Getenv(name)
	if p == "" {
		return nil, errNotFound
	}
	return []byte(p), nil
}
