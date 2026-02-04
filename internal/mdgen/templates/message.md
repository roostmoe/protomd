<h2 id="M_{{ .Name }}">Message <code>{{ .Name }}</code></h2>

{{ .Description }}

| Fields | |
|---|---|
{{ range $field := .Fields -}}
|`{{ $field.Name }}`|{{ $field.Description | prepInline }}|
{{ end -}}
