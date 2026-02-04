<h2 id="E_{{ .Name }}">Enum <code>{{ .Name }}</code></h2>

{{ .Description }}

| Options | |
|---|---|
{{ range $value := .Values -}}
|`{{ $value.Name }}`|{{ $value.Description | prepInline }}|
{{ end -}}
