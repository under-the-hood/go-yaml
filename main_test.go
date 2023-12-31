package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	yaml1 "github.com/goccy/go-yaml"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml2 "gopkg.in/yaml.v3"
)

func TestGoccyYamlParser(t *testing.T) {
	f, err := os.Open("testdata/typical.yaml")
	require.NoError(t, err)

	raw, err := io.ReadAll(f)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	t.Run("non-deterministic", func(t *testing.T) {
		var props any
		in := bytes.NewBuffer(raw)
		dec := yaml1.NewDecoder(in)
		assert.NoError(t, dec.Decode(&props))

		out := bytes.NewBuffer(nil)
		enc := yaml1.NewEncoder(out,
			yaml1.Indent(2),
			yaml1.IndentSequence(true),
		)
		assert.NoError(t, enc.Encode(props))
		assert.NotEqual(t, string(raw), out.String())
		// it should be equal, but has problems:
		// - ordering
		// - empty vs null
	})

	t.Run("deterministic", func(t *testing.T) {
		var props any
		in := bytes.NewBuffer(raw)
		dec := yaml1.NewDecoder(in, yaml1.UseOrderedMap())
		assert.NoError(t, dec.Decode(&props))

		out := bytes.NewBuffer(nil)
		enc := yaml1.NewEncoder(out,
			yaml1.Indent(2),
			yaml1.IndentSequence(true),
		)
		assert.NoError(t, enc.Encode(props))
		assert.NotEqual(t, string(raw), out.String())
		// it should be equal, but has problems:
		// - empty vs null
	})
}

func TestGopkgYamlParser(t *testing.T) {
	f, err := os.Open("testdata/typical.yaml")
	require.NoError(t, err)

	raw, err := io.ReadAll(f)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	t.Run("non-deterministic", func(t *testing.T) {
		var props any
		in := bytes.NewBuffer(raw)
		dec := yaml2.NewDecoder(in)
		assert.NoError(t, dec.Decode(&props))

		out := bytes.NewBuffer(nil)
		enc := yaml2.NewEncoder(out)
		enc.SetIndent(2)
		assert.NoError(t, enc.Encode(props))
		assert.NotEqual(t, string(raw), out.String())
		// it should be equal, but has problems:
		// - ordering
		// - empty vs null
	})

	t.Run("deterministic", func(t *testing.T) {
		var props yaml2.Node
		in := bytes.NewBuffer(raw)
		dec := yaml2.NewDecoder(in)
		assert.NoError(t, dec.Decode(&props))

		out := bytes.NewBuffer(nil)
		enc := yaml2.NewEncoder(out)
		enc.SetIndent(2)
		assert.NoError(t, enc.Encode(&props))
		assert.Equal(t, string(raw), out.String())
	})
}

func TestProperties(t *testing.T) {
	f, err := os.Open("testdata/typical.yaml")
	require.NoError(t, err)

	raw, err := io.ReadAll(f)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	t.Run("goccy yaml", func(t *testing.T) {
		var props PropsVar1
		in := bytes.NewBuffer(raw)
		dec := yaml1.NewDecoder(in)
		assert.NoError(t, dec.Decode(&props))

		props.Set("uid", uuid.New())

		out := bytes.NewBuffer(nil)
		enc := yaml1.NewEncoder(out,
			yaml1.Indent(2),
			yaml1.IndentSequence(true),
		)
		assert.NoError(t, enc.Encode(&props))
		assert.NotEqual(t, string(raw), out.String())
		// it should not be equal only by uid, but has another problems:
		// - empty vs null
	})

	t.Run("gopkg yaml", func(t *testing.T) {
		var props PropsVar2
		in := bytes.NewBuffer(raw)
		dec := yaml2.NewDecoder(in)
		assert.NoError(t, dec.Decode(&props))

		props.Set("uid", uuid.New())

		out := bytes.NewBuffer(nil)
		enc := yaml2.NewEncoder(out)
		enc.SetIndent(2)
		assert.NoError(t, enc.Encode(&props))
		assert.NotEqual(t, string(raw), out.String())
		// it should not be equal only by uid, but has another problems:
		// - ordering
		// - empty vs null
	})
}
