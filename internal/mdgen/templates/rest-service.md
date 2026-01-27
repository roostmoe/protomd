# {{ .Service.Name }} REST API Reference

{{ .Service.DocComment }}

## Methods

{{ range .Service.Methods }}
- [`{{ .Name }}`](./{{ .Name }}.md): {{ .DocComment | replace "\n\n" "<br>" | replace "\n" " " }}
{{ end }}
