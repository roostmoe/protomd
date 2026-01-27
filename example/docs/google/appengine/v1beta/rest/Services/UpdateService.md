# UpdateService

Updates the configuration of the specified service.

## HTTP request

```
PATCH /v1beta/{name=apps/*/services/*}
```

This URL uses [gRPC Transcoding](https://google.aip.dev/127) syntax.


## Request body
The request body contains an instance of [`UpdateServiceRequest`](../../index.md#UpdateServiceRequest).


## Response body
If successful, the response body contains an instance of [`Operation`](../../index.md#Operation).
