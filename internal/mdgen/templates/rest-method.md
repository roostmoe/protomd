# {{ .Method.Name }}

{{ .Method.DocComment }}

## HTTP request

```
{{ .Method.RestMethod }} {{ .Method.RestPath }}
```

This URL uses [gRPC Transcoding](https://google.aip.dev/127) syntax.

{{ if not (eq .Method.RestMethod "GET") }}
## Request body
The request body contains an instance of [`{{ .Method.Input.Name }}`]({{ namespaceMessageLink .Method.Input.Name "../.." }}).
{{ end }}

## Response body
If successful, the response body contains an instance of [`{{ .Method.Output.Name }}`]({{ namespaceMessageLink .Method.Output.Name "../.." }}).
