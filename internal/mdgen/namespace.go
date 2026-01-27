package mdgen

import (
	"fmt"
	"io"
	"strings"

	md "github.com/nao1215/markdown"
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

func (n Namespace) Build(w io.Writer) error {
	builder := md.NewMarkdown(w).
		H1f("Package %v", n.Name).
		H2("Index")

	var indexItems []string

	for _, s := range n.Services {
		indexItems = append(indexItems, fmt.Sprintf(
			"Service %s (RPC)",
			md.Link(s.Name, fmt.Sprintf("./rpc/%v/index.md", s.Name)),
		))

		if s.HasHttp {
			indexItems = append(indexItems, fmt.Sprintf(
				"Service %s (HTTP)",
				md.Link(s.Name, fmt.Sprintf("./rest/%v/index.md", s.Name)),
			))
		}
	}

	for _, e := range n.Enums {
		indexItems = append(indexItems, fmt.Sprintf(
			"Enum %s",
			md.Link(e.Name, fmt.Sprintf("#enum-%v", strings.ToLower(e.Name))),
		))
	}

	for _, m := range n.Messages {
		indexItems = append(indexItems, fmt.Sprintf(
			"Message %s",
			md.Link(m.Name, fmt.Sprintf("#message-%v", strings.ToLower(m.Name))),
		))
	}

	builder = builder.
		BulletList(indexItems...).
		HorizontalRule()

	for _, e := range n.Enums {
		builder = e.Build(builder)
	}

	for _, m := range n.Messages {
		builder = m.Build(builder)
	}

	return builder.Build()
}

func (n Namespace) Write(w io.Writer) error {
	type data struct {
		Namespace
	}

	return tmpl.ExecuteTemplate(w, "namespace.md", data{n})
}
