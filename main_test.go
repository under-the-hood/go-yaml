package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	yaml2 "github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "gopkg.in/yaml.v3"
)

func TestGoccyYamlParser(t *testing.T) {
	f, err := os.Open("testdata/typical.yaml")
	require.NoError(t, err)

	raw, err := io.ReadAll(f)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	t.Run("deterministic", func(t *testing.T) {
		var props any
		in := bytes.NewBuffer(raw)
		dec := yaml2.NewDecoder(in, yaml2.UseOrderedMap())
		assert.NoError(t, dec.Decode(&props))

		out := bytes.NewBuffer(nil)
		enc := yaml2.NewEncoder(out, yaml2.Indent(2))
		assert.NoError(t, enc.Encode(props))
		assert.NotEqual(t, string(raw), out.String())
		// problems:
		// - indent
		// - empty vs null
	})
}
