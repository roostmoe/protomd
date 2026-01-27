# ListInstances

Lists the instances of a version.

 Tip: To aggregate details about instances over time, see the
 [Stackdriver Monitoring API](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list).

## HTTP request

```
GET /v1beta/{parent=apps/*/services/*/versions/*}/instances
```

This URL uses [gRPC Transcoding](https://google.aip.dev/127) syntax.



## Response body
If successful, the response body contains an instance of [`ListInstancesResponse`](../../index.md#ListInstancesResponse).
