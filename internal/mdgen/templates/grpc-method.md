# {{ .Method.Name }}

{{ .Method.DocComment }}

## gRPC request

```
rpc {{ .Method.Name }}({{ .Method.Input.Name }}) returns ({{ .Method.Output.Name }}) {}
```

{{ if not (eq .Method.RestMethod "GET") }}
## Request body
The request body contains an instance of [`{{ .Method.Input.Name }}`]({{ namespaceMessageLink .Method.Input.Name "../.." }}).
{{ end }}

## Response body
If successful, the response body contains an instance of [`{{ .Method.Output.Name }}`]({{ namespaceMessageLink .Method.Output.Name "../.." }}).
