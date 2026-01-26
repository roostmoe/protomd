# {{ .Service.Name }} REST API Reference

{{ .Service.DocComment }}

## Methods

{{ range .Service.Methods }}
- [`{{ .Name }}`](./{{ .Name }}.md): {{ .DocComment }}
{{ end }}
