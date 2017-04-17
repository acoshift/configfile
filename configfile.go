package configfile

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
)

// Reader is the config reader
type Reader struct {
	base string
}

// NewReader creates new config reader with custom base path
func NewReader(base string) *Reader {
	return &Reader{base: base}
}

func (r *Reader) read(name string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(r.base, name))
}

// BytesDefault reads bytes from config file with default value
func (r *Reader) BytesDefault(name string, def []byte) []byte {
	b, err := r.read(name)
	if err != nil {
		return def
	}
	return b
}

// Bytes reads bytes from config file
func (r *Reader) Bytes(name string) []byte {
	return r.BytesDefault(name, nil)
}

// StringDefault reads string from config file with default value
func (r *Reader) StringDefault(name string, def string) string {
	b, err := r.read(name)
	if err != nil {
		return def
	}
	return string(b)
}

// String reads string from config file
func (r *Reader) String(name string) string {
	return r.StringDefault(name, "")
}

// IntDefault reads int from config file with default value
func (r *Reader) IntDefault(name string, def int) int {
	b, err := r.read(name)
	if err != nil {
		return def
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return def
	}
	return i
}

// Int reads int from config file
func (r *Reader) Int(name string) int {
	return r.IntDefault(name, 0)
}
