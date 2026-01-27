# Package {{ .Name }}

## Index

{{ range .Services -}}
- Service [{{ .Name }} (gRPC)](./rpc/{{ .Name }}/index.md)
{{- if .HasHttp }}
- Service [{{ .Name }} (REST)](./rest/{{ .Name }}/index.md)
{{ end -}}
{{ end -}}
{{ range .Enums -}}
- Enum [{{ .Name }}](#{{ .Name }})
{{ end -}}
{{ range .Messages -}}
- Message [{{ .Name }}](#{{ .Name }})
{{ end -}}

---

{{ range .Enums }}
<h2 id="{{ .Name }}">Enum <code>{{ .Name }}</code></h2>

{{ .DocComment }}

| Options | |
|---|---|
{{ range $option := .Options -}}
|`{{ $option.Name }}`|{{ $option.DocComment | replace "\n" "<br>" }}|
{{ end -}}
{{- end }}

{{ range .Messages }}
<h2 id="{{ .Name }}">Message <code>{{ .Name }}</code></h2>

{{ .DocComment }}

| Fields | |
|---|---|
{{ range $field := .Fields -}}
|`{{ $field.Name }}`|**{{ if $field.IsWellKnown }}[{{ $field.FieldType }}]({{ $field.WellKnownURL }}){{ else if $field.LinkedType }}[{{ $field.FieldType }}]({{ namespaceMessageLink $field.LinkedType "." }}){{ else }}{{ $field.FieldType }}{{ end }}**<br><br>{{ $field.DocComment | replace "\n" "<br>" }}|
{{ end -}}
{{- range $oneOf := .Oneofs -}}
|`{{ $oneOf.Name }}`|**oneOf**<br><br>{{ $oneOf.DocComment }}<br><hr/>Can be one of:<br><ul>{{ range $of := $oneOf.Fields }}<li>{{ $of.Name }} (<b>{{ if $of.IsWellKnown }}[{{ $of.FieldType }}]({{ $of.WellKnownURL }}){{ else if $of.LinkedType }}[{{ $of.FieldType }}]({{ namespaceMessageLink $of.LinkedType "." }}){{ else }}{{ $of.FieldType }}{{ end }}</b>)<br>{{ $of.DocComment | replace "\n" "<br>" }}</li><br>{{ end }}</ul>|
{{ end }}
{{- end }}
