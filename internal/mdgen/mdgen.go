package mdgen

import (
	"embed"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

//go:embed templates/*
var templateFs embed.FS
var tmpl *template.Template

// namespaceMessageLink generates a link to a message in a namespace's index
// accepts optional pathPrefix (defaults to "..")
func namespaceMessageLink(messageName string, pathPrefix ...string) string {
	prefix := ".."
	if len(pathPrefix) > 0 && pathPrefix[0] != "" {
		prefix = pathPrefix[0]
	}
	return prefix + "/index.md#" + messageName
}

func init() {
	funcMap := sprig.FuncMap()
	funcMap["namespaceMessageLink"] = namespaceMessageLink

	tmpl = template.Must(
		template.
			New("md").
			Funcs(funcMap).
			ParseFS(templateFs, "templates/*"),
	)
}
