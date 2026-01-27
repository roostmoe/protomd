# Services gRPC Reference

Manages services of an application.

## Methods


|ListServices|
|---|
|`rpc ListServices(`[`ListServicesRequest`](../../#ListServicesRequest)`) returns (`[`ListServicesResponse`](../../#ListServicesResponse)`)`<br><br>Lists all the services in the application.|

|GetService|
|---|
|`rpc GetService(`[`GetServiceRequest`](../../#GetServiceRequest)`) returns (`[`Service`](../../#Service)`)`<br><br>Gets the current configuration of the specified service.|

|UpdateService|
|---|
|`rpc UpdateService(`[`UpdateServiceRequest`](../../#UpdateServiceRequest)`) returns (`[`Operation`](../../#Operation)`)`<br><br>Updates the configuration of the specified service.|

|DeleteService|
|---|
|`rpc DeleteService(`[`DeleteServiceRequest`](../../#DeleteServiceRequest)`) returns (`[`Operation`](../../#Operation)`)`<br><br>Deletes the specified service and all enclosed versions.|

