package configfile_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/acoshift/configfile"
)

func testReader(t *testing.T, c *configfile.Reader) {
	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		t.Run("Bool", func(t *testing.T) {
			t.Parallel()

			assert.False(t, c.Bool("notfound"))
			assert.False(t, c.BoolDefault("notfound", false))
			assert.True(t, c.BoolDefault("notfound", true))
			assert.Panics(t, func() { c.MustBool("notfound") })
		})

		t.Run("Int", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 0, c.Int("notfound"))
			assert.Equal(t, 0, c.IntDefault("notfound", 0))
			assert.Equal(t, 1, c.IntDefault("notfound", 1))
			assert.Panics(t, func() { c.MustInt("notfound") })
		})

		t.Run("Int64", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, int64(0), c.Int64("notfound"))
			assert.Equal(t, int64(0), c.Int64Default("notfound", 0))
			assert.Equal(t, int64(1), c.Int64Default("notfound", 1))
			assert.Panics(t, func() { c.MustInt64("notfound") })
		})

		t.Run("Duration", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, time.Duration(0), c.Duration("notfound"))
			assert.Equal(t, time.Duration(0), c.DurationDefault("notfound", 0))
			assert.Equal(t, time.Duration(1), c.DurationDefault("notfound", 1))
			assert.Panics(t, func() { c.MustDuration("notfound") })
		})

		t.Run("String", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "", c.String("notfound"))
			assert.Equal(t, "", c.StringDefault("notfound", ""))
			assert.Equal(t, "a string", c.StringDefault("notfound", "a string"))
			assert.Panics(t, func() { c.MustString("notfound") })
		})

		t.Run("Bytes", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, []byte{}, c.Bytes("notfound"))
			assert.Nil(t, c.BytesDefault("notfound", nil))
			assert.Equal(t, []byte{}, c.BytesDefault("notfound", []byte{}))
			assert.Equal(t, []byte("some bytes"), c.BytesDefault("notfound", []byte("some bytes")))
			assert.Panics(t, func() { c.MustBytes("notfound") })
		})

		t.Run("Base64", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, []byte(""), c.Base64("notfound"))
			assert.Empty(t, c.Base64Default("notfound", nil))
			assert.Equal(t, []byte(""), c.Base64Default("notfound", []byte{}))
			assert.Equal(t, []byte("some bytes"), c.Base64Default("notfound", []byte("some bytes")))
			assert.Panics(t, func() { c.MustBase64("notfound") })
		})
	})

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

		t.Run("Int64", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, int64(0), c.Int64("empty"))
			assert.Equal(t, int64(0), c.Int64Default("empty", 0))
			assert.Equal(t, int64(1), c.Int64Default("empty", 1))
			assert.Panics(t, func() { c.MustInt64("empty") })
		})

		t.Run("Duration", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, time.Duration(0), c.Duration("empty"))
			assert.Equal(t, time.Duration(0), c.DurationDefault("empty", 0))
			assert.Equal(t, time.Duration(1), c.DurationDefault("empty", 1))
			assert.Panics(t, func() { c.MustDuration("empty") })
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

		t.Run("Base64", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, []byte(""), c.Base64("notfound"))
			assert.Empty(t, nil, c.Base64Default("notfound", nil))
			assert.Equal(t, []byte(""), c.Base64Default("notfound", []byte{}))
			assert.Equal(t, []byte("some bytes"), c.Base64Default("notfound", []byte("some bytes")))
			assert.Panics(t, func() { c.MustBase64("notfound") })
		})
	})

	t.Run("Data1", func(t *testing.T) {
		t.Parallel()

		t.Run("Bool", func(t *testing.T) {
			t.Parallel()

			assert.True(t, c.Bool("data1"))
			assert.True(t, c.BoolDefault("data1", false))
			assert.True(t, c.BoolDefault("data1", true))
			assert.NotPanics(t, func() { c.MustBool("data1") })
		})

		t.Run("Int", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 0, c.Int("data1"))
			assert.Equal(t, 0, c.IntDefault("data1", 0))
			assert.Equal(t, 1, c.IntDefault("data1", 1))
			assert.Panics(t, func() { c.MustInt("data1") })
		})

		t.Run("Int64", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, int64(0), c.Int64("data1"))
			assert.Equal(t, int64(0), c.Int64Default("data1", 0))
			assert.Equal(t, int64(1), c.Int64Default("data1", 1))
			assert.Panics(t, func() { c.MustInt64("data1") })
		})

		t.Run("Duration", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, time.Duration(0), c.Duration("data1"))
			assert.Equal(t, time.Duration(0), c.DurationDefault("data1", 0))
			assert.Equal(t, time.Duration(1), c.DurationDefault("data1", 1))
			assert.Panics(t, func() { c.MustDuration("data1") })
		})

		t.Run("String", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "true", c.String("data1"))
			assert.Equal(t, "true", c.StringDefault("data1", ""))
			assert.Equal(t, "true", c.StringDefault("data1", "a string"))
			assert.NotPanics(t, func() { c.MustString("data1") })
		})

		t.Run("Bytes", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, []byte("true"), c.Bytes("data1"))
			assert.Equal(t, []byte("true"), c.BytesDefault("data1", nil))
			assert.Equal(t, []byte("true"), c.BytesDefault("data1", []byte{}))
			assert.Equal(t, []byte("true"), c.BytesDefault("data1", []byte("some bytes")))
			assert.NotPanics(t, func() { c.MustBytes("data1") })
		})
	})

	t.Run("Data2", func(t *testing.T) {
		t.Parallel()

		t.Run("Bool", func(t *testing.T) {
			t.Parallel()

			assert.False(t, c.Bool("data2"))
			assert.False(t, c.BoolDefault("data2", false))
			assert.False(t, c.BoolDefault("data2", true))
			assert.NotPanics(t, func() { c.MustBool("data2") })
		})

		t.Run("Int", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 0, c.Int("data2"))
			assert.Equal(t, 0, c.IntDefault("data2", 0))
			assert.Equal(t, 1, c.IntDefault("data2", 1))
			assert.Panics(t, func() { c.MustInt("data2") })
		})

		t.Run("Int64", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, int64(0), c.Int64("data2"))
			assert.Equal(t, int64(0), c.Int64Default("data2", 0))
			assert.Equal(t, int64(1), c.Int64Default("data2", 1))
			assert.Panics(t, func() { c.MustInt64("data2") })
		})

		t.Run("Duration", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, time.Duration(0), c.Duration("data2"))
			assert.Equal(t, time.Duration(0), c.DurationDefault("data2", 0))
			assert.Equal(t, time.Duration(1), c.DurationDefault("data2", 1))
			assert.Panics(t, func() { c.MustDuration("data2") })
		})

		t.Run("String", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "false", c.String("data2"))
			assert.Equal(t, "false", c.StringDefault("data2", ""))
			assert.Equal(t, "false", c.StringDefault("data2", "a string"))
			assert.NotPanics(t, func() { c.MustString("data2") })
		})

		t.Run("Bytes", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, []byte("false"), c.Bytes("data2"))
			assert.Equal(t, []byte("false"), c.BytesDefault("data2", nil))
			assert.Equal(t, []byte("false"), c.BytesDefault("data2", []byte{}))
			assert.Equal(t, []byte("false"), c.BytesDefault("data2", []byte("some bytes")))
			assert.NotPanics(t, func() { c.MustBytes("data2") })
		})
	})

	t.Run("Data3", func(t *testing.T) {
		t.Parallel()

		t.Run("Bool", func(t *testing.T) {
			t.Parallel()

			assert.True(t, c.Bool("data3"))
			assert.True(t, c.BoolDefault("data3", false))
			assert.True(t, c.BoolDefault("data3", true))
			assert.NotPanics(t, func() { c.MustBool("data3") })
		})

		t.Run("Int", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 9, c.Int("data3"))
			assert.Equal(t, 9, c.IntDefault("data3", 0))
			assert.Equal(t, 9, c.IntDefault("data3", 1))
			assert.NotPanics(t, func() { c.MustInt("data3") })
		})

		t.Run("Int64", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, int64(9), c.Int64("data3"))
			assert.Equal(t, int64(9), c.Int64Default("data3", 0))
			assert.Equal(t, int64(9), c.Int64Default("data3", 1))
			assert.NotPanics(t, func() { c.MustInt64("data3") })
		})

		t.Run("Duration", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, time.Duration(0), c.Duration("data3"))
			assert.Equal(t, time.Duration(0), c.DurationDefault("data3", 0))
			assert.Equal(t, time.Duration(1), c.DurationDefault("data3", 1))
			assert.Panics(t, func() { c.MustDuration("data3") })
		})

		t.Run("String", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "9", c.String("data3"))
			assert.Equal(t, "9", c.StringDefault("data3", ""))
			assert.Equal(t, "9", c.StringDefault("data3", "a string"))
			assert.NotPanics(t, func() { c.MustString("data3") })
		})

		t.Run("Bytes", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, []byte("9"), c.Bytes("data3"))
			assert.Equal(t, []byte("9"), c.BytesDefault("data3", nil))
			assert.Equal(t, []byte("9"), c.BytesDefault("data3", []byte{}))
			assert.Equal(t, []byte("9"), c.BytesDefault("data3", []byte("some bytes")))
			assert.NotPanics(t, func() { c.MustBytes("data3") })
		})
	})

	t.Run("Data4", func(t *testing.T) {
		t.Parallel()

		t.Run("Bool", func(t *testing.T) {
			t.Parallel()

			assert.False(t, c.Bool("data4"))
			assert.False(t, c.BoolDefault("data4", false))
			assert.False(t, c.BoolDefault("data4", true))
			assert.NotPanics(t, func() { c.MustBool("data4") })
		})

		t.Run("Int", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 0, c.Int("data4"))
			assert.Equal(t, 0, c.IntDefault("data4", 0))
			assert.Equal(t, 0, c.IntDefault("data4", 1))
			assert.NotPanics(t, func() { c.MustInt("data4") })
		})

		t.Run("Int64", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, int64(0), c.Int64("data4"))
			assert.Equal(t, int64(0), c.Int64Default("data4", 0))
			assert.Equal(t, int64(0), c.Int64Default("data4", 1))
			assert.NotPanics(t, func() { c.MustInt64("data4") })
		})

		t.Run("Duration", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, time.Duration(0), c.Duration("data4"))
			assert.Equal(t, time.Duration(0), c.DurationDefault("data4", 0))
			assert.Equal(t, time.Duration(0), c.DurationDefault("data4", 1))
			assert.NotPanics(t, func() { c.MustDuration("data4") })
		})

		t.Run("String", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "0", c.String("data4"))
			assert.Equal(t, "0", c.StringDefault("data4", ""))
			assert.Equal(t, "0", c.StringDefault("data4", "a string"))
			assert.NotPanics(t, func() { c.MustString("data4") })
		})

		t.Run("Bytes", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, []byte("0"), c.Bytes("data4"))
			assert.Equal(t, []byte("0"), c.BytesDefault("data4", nil))
			assert.Equal(t, []byte("0"), c.BytesDefault("data4", []byte{}))
			assert.Equal(t, []byte("0"), c.BytesDefault("data4", []byte("some bytes")))
			assert.NotPanics(t, func() { c.MustBytes("data4") })
		})
	})

	t.Run("Data5", func(t *testing.T) {
		t.Parallel()

		t.Run("Bool", func(t *testing.T) {
			t.Parallel()

			assert.True(t, c.Bool("data5"))
			assert.True(t, c.BoolDefault("data5", false))
			assert.True(t, c.BoolDefault("data5", true))
			assert.NotPanics(t, func() { c.MustBool("data5") })
		})

		t.Run("Int", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 0, c.Int("data5"))
			assert.Equal(t, 0, c.IntDefault("data5", 0))
			assert.Equal(t, 1, c.IntDefault("data5", 1))
			assert.Panics(t, func() { c.MustInt("data5") })
		})

		t.Run("Int64", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, int64(0), c.Int64("data5"))
			assert.Equal(t, int64(0), c.Int64Default("data5", 0))
			assert.Equal(t, int64(1), c.Int64Default("data5", 1))
			assert.Panics(t, func() { c.MustInt64("data5") })
		})

		t.Run("Duration", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 3*time.Minute+5*time.Second, c.Duration("data5"))
			assert.Equal(t, 3*time.Minute+5*time.Second, c.DurationDefault("data5", 0))
			assert.Equal(t, 3*time.Minute+5*time.Second, c.DurationDefault("data5", 1))
			assert.NotPanics(t, func() { c.MustDuration("data5") })
		})

		t.Run("String", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "3m5s", c.String("data5"))
			assert.Equal(t, "3m5s", c.StringDefault("data5", ""))
			assert.Equal(t, "3m5s", c.StringDefault("data5", "a string"))
			assert.NotPanics(t, func() { c.MustString("data5") })
		})

		t.Run("Bytes", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, []byte("3m5s"), c.Bytes("data5"))
			assert.Equal(t, []byte("3m5s"), c.BytesDefault("data5", nil))
			assert.Equal(t, []byte("3m5s"), c.BytesDefault("data5", []byte{}))
			assert.Equal(t, []byte("3m5s"), c.BytesDefault("data5", []byte("some bytes")))
			assert.NotPanics(t, func() { c.MustBytes("data5") })
		})
	})

	t.Run("Data6", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, []byte("hello"), c.Base64("data6"))
		assert.Equal(t, []byte("hello"), c.Base64Default("data6", nil))
		assert.Equal(t, []byte("hello"), c.Base64Default("data6", []byte{}))
		assert.Equal(t, []byte("hello"), c.Base64Default("data6", []byte("some bytes")))
		assert.NotPanics(t, func() { c.MustBase64("data6") })
	})
}

func TestDirReader(t *testing.T) {
	t.Parallel()

	testReader(t, configfile.NewDirReader("testdata"))
	testReader(t, configfile.NewReader("testdata"))
}

func TestYAMLReader(t *testing.T) {
	t.Parallel()

	testReader(t, configfile.NewYAMLReader("testdata/config.yaml"))
	testReader(t, configfile.NewReader("testdata/config.yaml"))
}

func TestEnvReader(t *testing.T) {
	t.Parallel()

	testReader(t, configfile.NewEnvReader())
	testReader(t, configfile.NewReader("notexists"))
}
