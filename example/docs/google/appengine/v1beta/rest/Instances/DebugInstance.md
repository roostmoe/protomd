# DebugInstance

Enables debugging on a VM instance. This allows you to use the SSH
 command to connect to the virtual machine where the instance lives.
 While in "debug mode", the instance continues to serve live traffic.
 You should delete the instance when you are done debugging and then
 allow the system to take over and determine if another instance
 should be started.

 Only applicable for instances in App Engine flexible environment.

## HTTP request

```
POST /v1beta/{name=apps/*/services/*/versions/*/instances/*}:debug
```

This URL uses [gRPC Transcoding](https://google.aip.dev/127) syntax.


## Request body
The request body contains an instance of [`DebugInstanceRequest`](../../index.md#DebugInstanceRequest).


## Response body
If successful, the response body contains an instance of [`Operation`](../../index.md#Operation).
