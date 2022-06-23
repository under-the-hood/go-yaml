package main

import yaml2 "gopkg.in/yaml.v3"

type PropsVar1 struct{}

type PropsVar2 struct {
	*yaml2.Node
	data map[string]any
}

func (p *PropsVar2) Get(key string) any {
	return p.data[key]
}

func (p *PropsVar2) Set(key string, value any) {
	p.data[key] = value
}

func (p *PropsVar2) UnmarshalYAML(value *yaml2.Node) error {
	p.Node = value
	return value.Decode(&p.data)
}

func (p *PropsVar2) MarshalYAML() (any, error) {
	return p.Node, p.Encode(p.data)
}
