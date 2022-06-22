package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	yaml1 "github.com/goccy/go-yaml"
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
		enc := yaml1.NewEncoder(out, yaml1.Indent(2))
		assert.NoError(t, enc.Encode(props))
		assert.NotEqual(t, string(raw), out.String())
		// it should be equal, but has problems:
		// - ordering
		// - indent
		// - empty vs null
	})

	t.Run("deterministic", func(t *testing.T) {
		var props any
		in := bytes.NewBuffer(raw)
		dec := yaml1.NewDecoder(in, yaml1.UseOrderedMap())
		assert.NoError(t, dec.Decode(&props))

		out := bytes.NewBuffer(nil)
		enc := yaml1.NewEncoder(out, yaml1.Indent(2))
		assert.NoError(t, enc.Encode(props))
		assert.NotEqual(t, string(raw), out.String())
		// it should be equal, but has problems:
		// - indent
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
}
