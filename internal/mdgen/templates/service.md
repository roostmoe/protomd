<h2 id="S_{{ .Name }}">Service <code>{{ .Name }}</code></h2>

{{ .Description }}

{{ range .Methods -}}
{{ template "method.md" . }}
{{ end -}}
