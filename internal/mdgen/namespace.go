package mdgen

import (
	"io"
)

type Namespace struct {
	Name     string
	Services []Service
	Enums    []Enum
	Messages []Message
}

func NewNamespace(name string, services []Service, enums []Enum, messages []Message) Namespace {
	return Namespace{
		Name:     name,
		Services: services,
		Enums:    enums,
		Messages: messages,
	}
}

func (n Namespace) Write(w io.Writer) error {
	type data struct {
		Namespace
	}

	return tmpl.ExecuteTemplate(w, "namespace.md", data{n})
}
