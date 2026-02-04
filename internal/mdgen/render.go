package mdgen

import (
	"embed"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/roostmoe/protomd/internal/models"
)

//go:embed templates/*
var templateFs embed.FS
var tmpl *template.Template

func init() {
	funcMap := sprig.FuncMap()
	funcMap["messageLink"] = namespaceMessageLink
	funcMap["enumLink"] = enumMessageLink
	funcMap["prepInline"] = inlineTableMarkdown

	tmpl = template.Must(
		template.
			New("md").
			Funcs(funcMap).
			ParseFS(templateFs, "templates/*"),
	)
}

// namespaceMessageLink generates a link to a message in a namespace's index
// accepts optional pathPrefix (defaults to "..")
func namespaceMessageLink(messageName string, pathPrefix ...string) string {
	prefix := ".."
	if len(pathPrefix) > 0 && pathPrefix[0] != "" {
		prefix = pathPrefix[0]
	}
	return prefix + "/index.md#M_" + messageName
}

// enumMessageLink generates a link to a enum in a namespace's index
// accepts optional pathPrefix (defaults to "..")
func enumMessageLink(enumName string, pathPrefix ...string) string {
	prefix := ".."
	if len(pathPrefix) > 0 && pathPrefix[0] != "" {
		prefix = pathPrefix[0]
	}
	return prefix + "/index.md#E_" + enumName
}

func inlineTableMarkdown(text string) string {
	lines := strings.Split(text, "\n")
	for idx, line := range lines {
		lines[idx] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "<br/>")
}

func Render(pkg models.Package) (string, error) {
	var b strings.Builder

	if err := tmpl.ExecuteTemplate(&b, "namespace.md", pkg); err != nil {
		return "", err
	}

	return b.String(), nil
}
