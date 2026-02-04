# Package google.appengine.v1beta

## Index

- [`Services`](#S_Services) (interface)
- [`Instances`](#S_Instances) (interface)
- [`Instance`](#M_Instance) (message)
- [`NetworkSettings`](#M_NetworkSettings) (message)
- [`Service`](#M_Service) (message)
- [`TrafficSplit`](#M_TrafficSplit) (message)
- [`ListServicesRequest`](#M_ListServicesRequest) (message)
- [`ListServicesResponse`](#M_ListServicesResponse) (message)
- [`GetServiceRequest`](#M_GetServiceRequest) (message)
- [`UpdateServiceRequest`](#M_UpdateServiceRequest) (message)
- [`DeleteServiceRequest`](#M_DeleteServiceRequest) (message)
- [`ListInstancesRequest`](#M_ListInstancesRequest) (message)
- [`ListInstancesResponse`](#M_ListInstancesResponse) (message)
- [`GetInstanceRequest`](#M_GetInstanceRequest) (message)
- [`DeleteInstanceRequest`](#M_DeleteInstanceRequest) (message)
- [`DebugInstanceRequest`](#M_DebugInstanceRequest) (message)
- [`ListIngressRulesRequest`](#M_ListIngressRulesRequest) (message)

---

<h2 id="S_Services">Service <code>Services</code></h2>

Request message for `Services.ListServices`.

| ListServices |
|---|
|`rpc ListServices(...) returns (...)`<br/><br/>Name of the parent Application resource. Example: `apps/myapp`.|


| GetService |
|---|
|`rpc GetService(...) returns (...)`<br/><br/>Maximum results to return per page.|


| UpdateService |
|---|
|`rpc UpdateService(...) returns (...)`<br/><br/>Continuation token for fetching the next page of results.|


| DeleteService |
|---|
|`rpc DeleteService(...) returns (...)`<br/><br/>|



<h2 id="S_Instances">Service <code>Instances</code></h2>

Response message for `Services.ListServices`.

| ListInstances |
|---|
|`rpc ListInstances(...) returns (...)`<br/><br/>The services belonging to the requested application.|


| GetInstance |
|---|
|`rpc GetInstance(...) returns (...)`<br/><br/>Continuation token for fetching the next page of results.|


| DeleteInstance |
|---|
|`rpc DeleteInstance(...) returns (...)`<br/><br/>|


| DebugInstance |
|---|
|`rpc DebugInstance(...) returns (...)`<br/><br/>|



<h2 id="M_Instance">Message <code>Instance</code></h2>

An Instance resource is the computing unit that App Engine uses to
 automatically scale an application.

| Fields | |
|---|---|
|`name`|Output only. Full path to the Instance resource in the API.<br/>Example: `apps/myapp/services/default/versions/v1/instances/instance-1`.|
|`id`|Output only. Relative name of the instance within the version.<br/>Example: `instance-1`.|
|`app_engine_release`|Output only. App Engine release this instance is running on.|
|`availability`|Output only. Availability of the instance.|
|`vm_name`|Output only. Name of the virtual machine where this instance lives. Only applicable<br/>for instances in App Engine flexible environment.|
|`vm_zone_name`|Output only. Zone where the virtual machine is located. Only applicable for instances<br/>in App Engine flexible environment.|
|`vm_id`|Output only. Virtual machine ID of this instance. Only applicable for instances in<br/>App Engine flexible environment.|
|`start_time`|Output only. Time that this instance was started.<br/><br/>@OutputOnly|
|`requests`|Output only. Number of requests since this instance was started.|
|`errors`|Output only. Number of errors since this instance was started.|
|`qps`|Output only. Average queries per second (QPS) over the last minute.|
|`average_latency`|Output only. Average latency (ms) over the last minute.|
|`memory_usage`|Output only. Total memory in use (bytes).|
|`vm_status`|Output only. Status of the virtual machine where this instance lives. Only applicable<br/>for instances in App Engine flexible environment.|
|`vm_debug_enabled`|Output only. Whether this instance is in debug mode. Only applicable for instances in<br/>App Engine flexible environment.|
|`vm_ip`|Output only. The IP address of this instance. Only applicable for instances in App<br/>Engine flexible environment.|
|`vm_liveness`|Output only. The liveness health check of this instance. Only applicable for instances<br/>in App Engine flexible environment.|

<h2 id="M_NetworkSettings">Message <code>NetworkSettings</code></h2>

A NetworkSettings resource is a container for ingress settings for a version
 or service.

| Fields | |
|---|---|
|`ingress_traffic_allowed`|The ingress settings for version or service.|

<h2 id="M_Service">Message <code>Service</code></h2>

A Service resource is a logical component of an application that can share
 state and communicate in a secure fashion with other services.
 For example, an application that handles customer requests might
 include separate services to handle tasks such as backend data
 analysis or API requests from mobile devices. Each service has a
 collection of versions that define a specific set of code used to
 implement the functionality of that service.

| Fields | |
|---|---|
|`name`|Full path to the Service resource in the API.<br/>Example: `apps/myapp/services/default`.<br/><br/>@OutputOnly|
|`id`|Relative name of the service within the application.<br/>Example: `default`.<br/><br/>@OutputOnly|
|`split`|Mapping that defines fractional HTTP traffic diversion to<br/>different versions within the service.|
|`network_settings`|Ingress settings for this service. Will apply to all versions.|

<h2 id="M_TrafficSplit">Message <code>TrafficSplit</code></h2>

Traffic routing configuration for versions within a single service. Traffic
 splits define how traffic directed to the service is assigned to versions.

| Fields | |
|---|---|
|`shard_by`|Mechanism used to determine which version a request is sent to.<br/>The traffic selection algorithm will<br/>be stable for either type until allocations are changed.|
|`allocations`|Mapping from version IDs within the service to fractional<br/>(0.000, 1] allocations of traffic for that version. Each version can<br/>be specified only once, but some versions in the service may not<br/>have any traffic allocation. Services that have traffic allocated<br/>cannot be deleted until either the service is deleted or<br/>their traffic allocation is removed. Allocations must sum to 1.<br/>Up to two decimal place precision is supported for IP-based splits and<br/>up to three decimal places is supported for cookie-based splits.|

<h2 id="M_ListServicesRequest">Message <code>ListServicesRequest</code></h2>

Request message for `Services.ListServices`.

| Fields | |
|---|---|
|`parent`|Name of the parent Application resource. Example: `apps/myapp`.|
|`page_size`|Maximum results to return per page.|
|`page_token`|Continuation token for fetching the next page of results.|

<h2 id="M_ListServicesResponse">Message <code>ListServicesResponse</code></h2>

Response message for `Services.ListServices`.

| Fields | |
|---|---|
|`services`|The services belonging to the requested application.|
|`next_page_token`|Continuation token for fetching the next page of results.|

<h2 id="M_GetServiceRequest">Message <code>GetServiceRequest</code></h2>

Request message for `Services.GetService`.

| Fields | |
|---|---|
|`name`|Name of the resource requested. Example: `apps/myapp/services/default`.|

<h2 id="M_UpdateServiceRequest">Message <code>UpdateServiceRequest</code></h2>

Request message for `Services.UpdateService`.

| Fields | |
|---|---|
|`name`|Name of the resource to update. Example: `apps/myapp/services/default`.|
|`service`|A Service resource containing the updated service. Only fields set in the<br/>field mask will be updated.|
|`update_mask`|Standard field mask for the set of fields to be updated.|
|`migrate_traffic`|Set to `true` to gradually shift traffic to one or more versions that you<br/>specify. By default, traffic is shifted immediately.<br/>For gradual traffic migration, the target versions<br/>must be located within instances that are configured for both<br/>[warmup requests](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1beta/apps.services.versions#InboundServiceType)<br/>and<br/>[automatic scaling](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1beta/apps.services.versions#AutomaticScaling).<br/>You must specify the<br/>[`shardBy`](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1beta/apps.services#ShardBy)<br/>field in the Service resource. Gradual traffic migration is not<br/>supported in the App Engine flexible environment. For examples, see<br/>[Migrating and Splitting Traffic](https://cloud.google.com/appengine/docs/admin-api/migrating-splitting-traffic).|

<h2 id="M_DeleteServiceRequest">Message <code>DeleteServiceRequest</code></h2>

Request message for `Services.DeleteService`.

| Fields | |
|---|---|
|`name`|Name of the resource requested. Example: `apps/myapp/services/default`.|

<h2 id="M_ListInstancesRequest">Message <code>ListInstancesRequest</code></h2>

Request message for `Instances.ListInstances`.

| Fields | |
|---|---|
|`parent`|Name of the parent Version resource. Example:<br/>`apps/myapp/services/default/versions/v1`.|
|`page_size`|Maximum results to return per page.|
|`page_token`|Continuation token for fetching the next page of results.|

<h2 id="M_ListInstancesResponse">Message <code>ListInstancesResponse</code></h2>

Response message for `Instances.ListInstances`.

| Fields | |
|---|---|
|`instances`|The instances belonging to the requested version.|
|`next_page_token`|Continuation token for fetching the next page of results.|

<h2 id="M_GetInstanceRequest">Message <code>GetInstanceRequest</code></h2>

Request message for `Instances.GetInstance`.

| Fields | |
|---|---|
|`name`|Name of the resource requested. Example:<br/>`apps/myapp/services/default/versions/v1/instances/instance-1`.|

<h2 id="M_DeleteInstanceRequest">Message <code>DeleteInstanceRequest</code></h2>

Request message for `Instances.DeleteInstance`.

| Fields | |
|---|---|
|`name`|Name of the resource requested. Example:<br/>`apps/myapp/services/default/versions/v1/instances/instance-1`.|

<h2 id="M_DebugInstanceRequest">Message <code>DebugInstanceRequest</code></h2>

Request message for `Instances.DebugInstance`.

| Fields | |
|---|---|
|`name`|Name of the resource requested. Example:<br/>`apps/myapp/services/default/versions/v1/instances/instance-1`.|
|`ssh_key`|Public SSH key to add to the instance. Examples:<br/><br/>* `[USERNAME]:ssh-rsa [KEY_VALUE] [USERNAME]`<br/>* `[USERNAME]:ssh-rsa [KEY_VALUE] google-ssh {"userName":"[USERNAME]","expireOn":"[EXPIRE_TIME]"}`<br/><br/>For more information, see<br/>[Adding and Removing SSH Keys](https://cloud.google.com/compute/docs/instances/adding-removing-ssh-keys).|

<h2 id="M_ListIngressRulesRequest">Message <code>ListIngressRulesRequest</code></h2>

Request message for `Firewall.ListIngressRules`.

| Fields | |
|---|---|
|`parent`|Name of the Firewall collection to retrieve.<br/>Example: `apps/myapp/firewall/ingressRules`.|
|`page_size`|Maximum results to return per page.|
|`page_token`|Continuation token for fetching the next page of results.|
|`matching_address`|A valid IP Address. If set, only rules matching this address will be<br/>returned. The first returned rule will be the rule that fires on requests<br/>from this IP.|

