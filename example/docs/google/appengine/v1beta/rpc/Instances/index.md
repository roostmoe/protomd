# Instances gRPC Reference

Manages instances of a version.

## Methods


|ListInstances|
|---|
|`rpc ListInstances(`[`ListInstancesRequest`](../../#ListInstancesRequest)`) returns (`[`ListInstancesResponse`](../../#ListInstancesResponse)`)`<br><br>Lists the instances of a version.

 Tip: To aggregate details about instances over time, see the
 [Stackdriver Monitoring API](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list).|

|GetInstance|
|---|
|`rpc GetInstance(`[`GetInstanceRequest`](../../#GetInstanceRequest)`) returns (`[`Instance`](../../#Instance)`)`<br><br>Gets instance information.|

|DeleteInstance|
|---|
|`rpc DeleteInstance(`[`DeleteInstanceRequest`](../../#DeleteInstanceRequest)`) returns (`[`Operation`](../../#Operation)`)`<br><br>Stops a running instance.

 The instance might be automatically recreated based on the scaling settings
 of the version. For more information, see "How Instances are Managed"
 ([standard environment](https://cloud.google.com/appengine/docs/standard/python/how-instances-are-managed) |
 [flexible environment](https://cloud.google.com/appengine/docs/flexible/python/how-instances-are-managed)).

 To ensure that instances are not re-created and avoid getting billed, you
 can stop all instances within the target version by changing the serving
 status of the version to `STOPPED` with the
 [`apps.services.versions.patch`](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions/patch)
 method.|

|DebugInstance|
|---|
|`rpc DebugInstance(`[`DebugInstanceRequest`](../../#DebugInstanceRequest)`) returns (`[`Operation`](../../#Operation)`)`<br><br>Enables debugging on a VM instance. This allows you to use the SSH
 command to connect to the virtual machine where the instance lives.
 While in "debug mode", the instance continues to serve live traffic.
 You should delete the instance when you are done debugging and then
 allow the system to take over and determine if another instance
 should be started.

 Only applicable for instances in App Engine flexible environment.|

