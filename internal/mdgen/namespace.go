package mdgen

import (
	"io"
)

type Namespace struct {
	Name     string
	Services []Service
	Messages []Message
}

func NewNamespace(name string, services []Service, messages []Message) Namespace {
	return Namespace{
		Name:     name,
		Services: services,
		Messages: messages,
	}
}

func (n Namespace) Write(w io.Writer) error {
	type data struct {
		Namespace
	}

	return tmpl.ExecuteTemplate(w, "namespace.md", data{n})
}
