# Package google.longrunning

## Index

- [`Operations`](#S_Operations) (interface)
- [`Operation`](#M_Operation) (message)
- [`GetOperationRequest`](#M_GetOperationRequest) (message)
- [`ListOperationsRequest`](#M_ListOperationsRequest) (message)
- [`ListOperationsResponse`](#M_ListOperationsResponse) (message)
- [`CancelOperationRequest`](#M_CancelOperationRequest) (message)
- [`DeleteOperationRequest`](#M_DeleteOperationRequest) (message)
- [`WaitOperationRequest`](#M_WaitOperationRequest) (message)
- [`OperationInfo`](#M_OperationInfo) (message)

---

<h2 id="S_Operations">Service <code>Operations</code></h2>

This resource represents a long-running operation that is the result of a
 network API call.

| ListOperations |
|---|
|`rpc ListOperations(...) returns (...)`<br/><br/>The server-assigned name, which is only unique within the same service that<br/>originally returns it. If you use the default HTTP mapping, the<br/>`name` should be a resource name ending with `operations/{unique_id}`.|


| GetOperation |
|---|
|`rpc GetOperation(...) returns (...)`<br/><br/>Service-specific metadata associated with the operation.  It typically<br/>contains progress information and common metadata such as create time.<br/>Some services might not provide such metadata.  Any method that returns a<br/>long-running operation should document the metadata type, if any.|


| DeleteOperation |
|---|
|`rpc DeleteOperation(...) returns (...)`<br/><br/>If the value is `false`, it means the operation is still in progress.<br/>If `true`, the operation is completed, and either `error` or `response` is<br/>available.|


| CancelOperation |
|---|
|`rpc CancelOperation(...) returns (...)`<br/><br/>The error result of the operation in case of failure or cancellation.|


| WaitOperation |
|---|
|`rpc WaitOperation(...) returns (...)`<br/><br/>The normal, successful response of the operation.  If the original<br/>method returns no data on success, such as `Delete`, the response is<br/>`google.protobuf.Empty`.  If the original method is standard<br/>`Get`/`Create`/`Update`, the response should be the resource.  For other<br/>methods, the response should have the type `XxxResponse`, where `Xxx`<br/>is the original method name.  For example, if the original method name<br/>is `TakeSnapshot()`, the inferred response type is<br/>`TakeSnapshotResponse`.|



<h2 id="M_Operation">Message <code>Operation</code></h2>

This resource represents a long-running operation that is the result of a
 network API call.

| Fields | |
|---|---|
|`name`|The server-assigned name, which is only unique within the same service that<br/>originally returns it. If you use the default HTTP mapping, the<br/>`name` should be a resource name ending with `operations/{unique_id}`.|
|`metadata`|Service-specific metadata associated with the operation.  It typically<br/>contains progress information and common metadata such as create time.<br/>Some services might not provide such metadata.  Any method that returns a<br/>long-running operation should document the metadata type, if any.|
|`done`|If the value is `false`, it means the operation is still in progress.<br/>If `true`, the operation is completed, and either `error` or `response` is<br/>available.|
|`error`|The error result of the operation in case of failure or cancellation.|
|`response`|The normal, successful response of the operation.  If the original<br/>method returns no data on success, such as `Delete`, the response is<br/>`google.protobuf.Empty`.  If the original method is standard<br/>`Get`/`Create`/`Update`, the response should be the resource.  For other<br/>methods, the response should have the type `XxxResponse`, where `Xxx`<br/>is the original method name.  For example, if the original method name<br/>is `TakeSnapshot()`, the inferred response type is<br/>`TakeSnapshotResponse`.|

<h2 id="M_GetOperationRequest">Message <code>GetOperationRequest</code></h2>

The request message for
 [Operations.GetOperation][google.longrunning.Operations.GetOperation].

| Fields | |
|---|---|
|`name`|The name of the operation resource.|

<h2 id="M_ListOperationsRequest">Message <code>ListOperationsRequest</code></h2>

The request message for
 [Operations.ListOperations][google.longrunning.Operations.ListOperations].

| Fields | |
|---|---|
|`name`|The name of the operation's parent resource.|
|`filter`|The standard list filter.|
|`page_size`|The standard list page size.|
|`page_token`|The standard list page token.|
|`return_partial_success`|When set to `true`, operations that are reachable are returned as normal,<br/>and those that are unreachable are returned in the<br/>[ListOperationsResponse.unreachable] field.<br/><br/>This can only be `true` when reading across collections e.g. when `parent`<br/>is set to `"projects/example/locations/-"`.<br/><br/>This field is not by default supported and will result in an<br/>`UNIMPLEMENTED` error if set unless explicitly documented otherwise in<br/>service or product specific documentation.|

<h2 id="M_ListOperationsResponse">Message <code>ListOperationsResponse</code></h2>

The response message for
 [Operations.ListOperations][google.longrunning.Operations.ListOperations].

| Fields | |
|---|---|
|`operations`|A list of operations that matches the specified filter in the request.|
|`next_page_token`|The standard List next-page token.|
|`unreachable`|Unordered list. Unreachable resources. Populated when the request sets<br/>`ListOperationsRequest.return_partial_success` and reads across<br/>collections e.g. when attempting to list all resources across all supported<br/>locations.|

<h2 id="M_CancelOperationRequest">Message <code>CancelOperationRequest</code></h2>

The request message for
 [Operations.CancelOperation][google.longrunning.Operations.CancelOperation].

| Fields | |
|---|---|
|`name`|The name of the operation resource to be cancelled.|

<h2 id="M_DeleteOperationRequest">Message <code>DeleteOperationRequest</code></h2>

The request message for
 [Operations.DeleteOperation][google.longrunning.Operations.DeleteOperation].

| Fields | |
|---|---|
|`name`|The name of the operation resource to be deleted.|

<h2 id="M_WaitOperationRequest">Message <code>WaitOperationRequest</code></h2>

The request message for
 [Operations.WaitOperation][google.longrunning.Operations.WaitOperation].

| Fields | |
|---|---|
|`name`|The name of the operation resource to wait on.|
|`timeout`|The maximum duration to wait before timing out. If left blank, the wait<br/>will be at most the time permitted by the underlying HTTP/RPC protocol.<br/>If RPC context deadline is also specified, the shorter one will be used.|

<h2 id="M_OperationInfo">Message <code>OperationInfo</code></h2>

A message representing the message types used by a long-running operation.

 Example:

     rpc Export(ExportRequest) returns (google.longrunning.Operation) {
       option (google.longrunning.operation_info) = {
         response_type: "ExportResponse"
         metadata_type: "ExportMetadata"
       };
     }

| Fields | |
|---|---|
|`response_type`|Required. The message name of the primary return type for this<br/>long-running operation.<br/>This type will be used to deserialize the LRO's response.<br/><br/>If the response is in a different package from the rpc, a fully-qualified<br/>message name must be used (e.g. `google.protobuf.Struct`).<br/><br/>Note: Altering this value constitutes a breaking change.|
|`metadata_type`|Required. The message name of the metadata type for this long-running<br/>operation.<br/><br/>If the response is in a different package from the rpc, a fully-qualified<br/>message name must be used (e.g. `google.protobuf.Struct`).<br/><br/>Note: Altering this value constitutes a breaking change.|

