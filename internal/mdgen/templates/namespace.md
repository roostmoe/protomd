# Package {{ .Name }}

## Index

{{ range .Services -}}
- [`{{ .Name }}`](#S_{{ .Name }}) (interface)
{{ end -}}
{{ range .Enums -}}
- [`{{ .Name }}`](#E_{{ .Name }}) (enum)
{{ end -}}
{{ range .Messages -}}
- [`{{ .Name }}`](#M_{{ .Name }}) (message)
{{ end }}
---

{{ range .Services -}}
{{ template "service.md" . }}
{{ end -}}

{{ range .Enums -}}
{{ template "enum.md" . }}
{{ end -}}

{{ range .Messages -}}
{{ template "message.md" . }}
{{ end -}}
