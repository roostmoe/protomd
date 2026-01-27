# DeleteInstance

Stops a running instance.

 The instance might be automatically recreated based on the scaling settings
 of the version. For more information, see "How Instances are Managed"
 ([standard environment](https://cloud.google.com/appengine/docs/standard/python/how-instances-are-managed) |
 [flexible environment](https://cloud.google.com/appengine/docs/flexible/python/how-instances-are-managed)).

 To ensure that instances are not re-created and avoid getting billed, you
 can stop all instances within the target version by changing the serving
 status of the version to `STOPPED` with the
 [`apps.services.versions.patch`](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions/patch)
 method.

## HTTP request

```
DELETE /v1beta/{name=apps/*/services/*/versions/*/instances/*}
```

This URL uses [gRPC Transcoding](https://google.aip.dev/127) syntax.


## Request body
The request body contains an instance of [`DeleteInstanceRequest`](../../index.md#DeleteInstanceRequest).


## Response body
If successful, the response body contains an instance of [`Operation`](../../index.md#Operation).
