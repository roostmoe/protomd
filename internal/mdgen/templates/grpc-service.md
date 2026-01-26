# {{ .Service.Name }} gRPC Reference

{{ .Service.DocComment }}

## Methods

{{ range .Service.Methods }}
|{{ .Name }}|
|---|
|`rpc {{ .Name }}(`[`{{ .Input.Name }}`](../../#{{ .Input.Name }})`) returns (`[`{{ .Output.Name }}`](../../#{{ .Output.Name }})`)`<br><br>{{ .DocComment }}|
{{ end }}
