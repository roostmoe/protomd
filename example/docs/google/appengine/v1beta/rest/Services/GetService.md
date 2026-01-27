# GetService

Gets the current configuration of the specified service.

## HTTP request

```
GET /v1beta/{name=apps/*/services/*}
```

This URL uses [gRPC Transcoding](https://google.aip.dev/127) syntax.



## Response body
If successful, the response body contains an instance of [`Service`](../../index.md#Service).
