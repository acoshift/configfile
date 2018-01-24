package configfile_test

import (
	"testing"

	"github.com/acoshift/configfile"
	"github.com/stretchr/testify/assert"
)

func TestConfigfile(t *testing.T) {
	c := configfile.NewReader("testdata")

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()

		t.Run("Bool", func(t *testing.T) {
			t.Parallel()

			assert.False(t, c.Bool("empty"))
			assert.False(t, c.BoolDefault("empty", false))
			assert.True(t, c.BoolDefault("empty", true))
			assert.Panics(t, func() { c.MustBool("empty") })
		})

		t.Run("Int", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 0, c.Int("empty"))
			assert.Equal(t, 0, c.IntDefault("empty", 0))
			assert.Equal(t, 1, c.IntDefault("empty", 1))
			assert.Panics(t, func() { c.MustInt("empty") })
		})

		t.Run("String", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "", c.String("empty"))
			assert.Equal(t, "", c.StringDefault("empty", ""))
			assert.Equal(t, "", c.StringDefault("empty", "a string"))
			assert.NotPanics(t, func() { c.MustString("empty") })
		})

		t.Run("Bytes", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, []byte{}, c.Bytes("empty"))
			assert.Equal(t, []byte{}, c.BytesDefault("empty", nil))
			assert.Equal(t, []byte{}, c.BytesDefault("empty", []byte{}))
			assert.Equal(t, []byte{}, c.BytesDefault("empty", []byte("some bytes")))
			assert.NotPanics(t, func() { c.MustBytes("empty") })
		})
	})
}
