package main

import (
	"bytes"

	yaml1 "github.com/goccy/go-yaml"
	yaml2 "gopkg.in/yaml.v3"
)

type PropsVar1 struct {
	yaml1.MapSlice
}

func (props *PropsVar1) Get(key string) any {
	for _, item := range props.MapSlice {
		if item.Key == key {
			return item.Value
		}
	}
	return nil
}

func (props *PropsVar1) Set(key string, value any) {
	for i, item := range props.MapSlice {
		if item.Key == key {
			props.MapSlice[i].Value = value
			return
		}
	}
	props.MapSlice = append(props.MapSlice, yaml1.MapItem{
		Key:   key,
		Value: value,
	})
	return
}

func (props *PropsVar1) Delete(key string) {
	for i, item := range props.MapSlice {
		if item.Key == key {
			props.MapSlice = append(props.MapSlice[:i], props.MapSlice[i+1:]...)
			return
		}
	}
	return
}

func (props *PropsVar1) UnmarshalYAML(val []byte) error {
	dec := yaml1.NewDecoder(bytes.NewBuffer(val))
	return dec.Decode(&props.MapSlice)
}

func (props *PropsVar1) MarshalYAML() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := yaml1.NewEncoder(buf, yaml1.Indent(2), yaml1.IndentSequence(true))
	err := enc.Encode(props.MapSlice)
	return buf.Bytes(), err
}

type PropsVar2 struct {
	*yaml2.Node
	data map[string]any
}

func (props *PropsVar2) Get(key string) any {
	if props.data == nil {
		return nil
	}
	return props.data[key]
}

func (props *PropsVar2) Set(key string, value any) {
	if props.data == nil {
		props.data = make(map[string]any)
	}
	props.data[key] = value
}

func (props *PropsVar2) UnmarshalYAML(val *yaml2.Node) error {
	props.Node = val
	return val.Decode(&props.data)
}

func (props *PropsVar2) MarshalYAML() (any, error) {
	err := props.Encode(props.data)
	return props.Node, err
}
