# Package {{ .Name }}

## Index

{{ range .Services -}}
{{ if .HasHttp -}}
- service [{{ .Name }} (REST)](./rest/{{ .Name }}/index.md)
{{ end -}}
- service [{{ .Name }} (gRPC)](./rpc/{{ .Name }}/index.md)
{{ end -}}
{{ range .Messages -}}
- message [{{ .Name }}](#{{ .Name }})
{{ end -}}

---

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
