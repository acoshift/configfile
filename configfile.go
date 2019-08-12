package configfile

import (
	"encoding/base64"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/acoshift/configfile/internal/reader"
)

// NewReader creates new config reader
func NewReader(base string) *Reader {
	stats, _ := os.Stat(base)
	if stats != nil {
		if stats.IsDir() {
			return NewDirReader(base)
		}
		return NewYAMLReader(base)
	}
	return NewEnvReader()
}

// NewDirReader creates new config dir reader
func NewDirReader(base string) *Reader {
	return &Reader{r: reader.NewDir(base)}
}

// NewYAMLReader creates new yaml reader
func NewYAMLReader(filename string) *Reader {
	return &Reader{r: reader.NewYAML(filename)}
}

// NewEnvReader creates new env reader
func NewEnvReader() *Reader {
	return &Reader{r: reader.NewEnv()}
}

type intlReader interface {
	Read(name string) ([]byte, error)
}

// Reader is the config reader
type Reader struct {
	r        intlReader
	fallback *Reader
}

func (r *Reader) Fallback(f *Reader) *Reader {
	r.fallback = f
	return r
}

func (r *Reader) read(name string) ([]byte, error) {
	b, err := r.r.Read(name)
	if err != nil && r.fallback != nil {
		b, err = r.fallback.read(name)
	}
	return b, err
}

func (r *Reader) readString(name string) (string, error) {
	b, err := r.read(name)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *Reader) readInt(name string) (int, error) {
	s, err := r.readString(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func (r *Reader) readInt64(name string) (int64, error) {
	s, err := r.readString(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(s, 10, 64)
}

func (r *Reader) readBool(name string) (bool, error) {
	s, err := r.readString(name)
	if err != nil {
		return false, err
	}
	if s == "" {
		return false, io.EOF
	}
	if s == "0" {
		return false, nil
	}
	if strings.ToLower(s) == "false" {
		return false, nil
	}
	return true, nil
}

func (r *Reader) readDuration(name string) (time.Duration, error) {
	s, err := r.readString(name)
	if err != nil {
		return 0, err
	}
	return time.ParseDuration(s)
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
	return r.BytesDefault(name, []byte{})
}

// MustBytes reads bytes from config file, panic if file not exists
func (r *Reader) MustBytes(name string) []byte {
	s, err := r.read(name)
	if err != nil {
		panic(err)
	}
	return s
}

// StringDefault reads string from config file with default value
func (r *Reader) StringDefault(name string, def string) string {
	s, err := r.readString(name)
	if err != nil {
		return def
	}
	return s
}

// String reads string from config file
func (r *Reader) String(name string) string {
	return r.StringDefault(name, "")
}

// MustString reads string from config file, panic if file not exists
func (r *Reader) MustString(name string) string {
	s, err := r.readString(name)
	if err != nil {
		panic(err)
	}
	return s
}

// Base64Default reads string from config file then decode using base64
// if error, will return default value
func (r *Reader) Base64Default(name string, def []byte) []byte {
	s, err := r.readString(name)
	if err != nil {
		return def
	}
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return def
	}
	return b
}

// Base64 reads string from config file then decode using base64
func (r *Reader) Base64(name string) []byte {
	return r.Base64Default(name, []byte{})
}

// MustBase64 reads string from config file then decode using base64
// if error, will panic
func (r *Reader) MustBase64(name string) []byte {
	s, err := r.readString(name)
	if err != nil {
		panic(err)
	}
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// IntDefault reads int from config file with default value
func (r *Reader) IntDefault(name string, def int) int {
	i, err := r.readInt(name)
	if err != nil {
		return def
	}
	return i
}

// Int reads int from config file
func (r *Reader) Int(name string) int {
	return r.IntDefault(name, 0)
}

// MustInt reads int from config file, panic if file not exists or data can not parse to int
func (r *Reader) MustInt(name string) int {
	i, err := r.readInt(name)
	if err != nil {
		panic(err)
	}
	return i
}

// Int64Default reads int64 from config file with default value
func (r *Reader) Int64Default(name string, def int64) int64 {
	i, err := r.readInt64(name)
	if err != nil {
		return def
	}
	return i
}

// Int64 reads int from config file
func (r *Reader) Int64(name string) int64 {
	return r.Int64Default(name, 0)
}

// MustInt64 reads int64 from config file, panic if file not exists or data can not parse to int64
func (r *Reader) MustInt64(name string) int64 {
	i, err := r.readInt64(name)
	if err != nil {
		panic(err)
	}
	return i
}

// BoolDefault reads bool from config file with default value,
// result is false if lower case data is "", "0", or "false", otherwise true
func (r *Reader) BoolDefault(name string, def bool) bool {
	b, err := r.readBool(name)
	if err != nil {
		return def
	}
	return b
}

// Bool reads bool from config file, see BoolDefault
func (r *Reader) Bool(name string) bool {
	return r.BoolDefault(name, false)
}

// MustBool reads bool from config file, see BoolDefault,
// panic if file not exists
func (r *Reader) MustBool(name string) bool {
	b, err := r.readBool(name)
	if err != nil {
		panic(err)
	}
	return b
}

// DurationDefault reads string then parse as duration from config file with default value
func (r *Reader) DurationDefault(name string, def time.Duration) time.Duration {
	d, err := r.readDuration(name)
	if err != nil {
		return def
	}
	return d
}

// Duration reads string then parse as duration from config file
func (r *Reader) Duration(name string) time.Duration {
	return r.DurationDefault(name, 0)
}

// MustDuration reads string then parse as duration from config file,
// panic if file not exists
func (r *Reader) MustDuration(name string) time.Duration {
	b, err := r.readDuration(name)
	if err != nil {
		panic(err)
	}
	return b
}
