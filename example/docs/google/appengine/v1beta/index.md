# Package google.appengine.v1beta

## Index

- Service [Instances (gRPC)](./rpc/Instances/index.md)
- Service [Instances (REST)](./rest/Instances/index.md)
- Service [Services (gRPC)](./rpc/Services/index.md)
- Service [Services (REST)](./rest/Services/index.md)
- Enum [Instance.Availability](#Instance.Availability)
- Enum [Availability](#Availability)
- Enum [Liveness.LivenessState](#Liveness.LivenessState)
- Enum [LivenessState](#LivenessState)
- Enum [TrafficSplit.ShardBy](#TrafficSplit.ShardBy)
- Enum [ShardBy](#ShardBy)
- Enum [Instance.Liveness.LivenessState](#Instance.Liveness.LivenessState)
- Enum [NetworkSettings.IngressTrafficAllowed](#NetworkSettings.IngressTrafficAllowed)
- Enum [IngressTrafficAllowed](#IngressTrafficAllowed)
- Message [Service](#Service)
- Message [TrafficSplit.AllocationsEntry](#TrafficSplit.AllocationsEntry)
- Message [ListServicesRequest](#ListServicesRequest)
- Message [ListServicesResponse](#ListServicesResponse)
- Message [GetServiceRequest](#GetServiceRequest)
- Message [DeleteServiceRequest](#DeleteServiceRequest)
- Message [ListInstancesRequest](#ListInstancesRequest)
- Message [ListInstancesResponse](#ListInstancesResponse)
- Message [Instance](#Instance)
- Message [Instance.Liveness](#Instance.Liveness)
- Message [Liveness](#Liveness)
- Message [GetInstanceRequest](#GetInstanceRequest)
- Message [NetworkSettings](#NetworkSettings)
- Message [ListIngressRulesRequest](#ListIngressRulesRequest)
- Message [TrafficSplit](#TrafficSplit)
- Message [AllocationsEntry](#AllocationsEntry)
- Message [UpdateServiceRequest](#UpdateServiceRequest)
- Message [DeleteInstanceRequest](#DeleteInstanceRequest)
- Message [DebugInstanceRequest](#DebugInstanceRequest)
---


<h2 id="Instance.Availability">Enum <code>Instance.Availability</code></h2>

Availability of the instance.

| Options | |
|---|---|
|`UNSPECIFIED`||
|`RESIDENT`||
|`DYNAMIC`||

<h2 id="Availability">Enum <code>Availability</code></h2>

Availability of the instance.

| Options | |
|---|---|
|`UNSPECIFIED`||
|`RESIDENT`||
|`DYNAMIC`||

<h2 id="Liveness.LivenessState">Enum <code>Liveness.LivenessState</code></h2>

Liveness health check status for Flex instances.

| Options | |
|---|---|
|`LIVENESS_STATE_UNSPECIFIED`|There is no liveness health check for the instance. Only applicable for<br> instances in App Engine standard environment.|
|`UNKNOWN`|The health checking system is aware of the instance but its health is<br> not known at the moment.|
|`HEALTHY`|The instance is reachable i.e. a connection to the application health<br> checking endpoint can be established, and conforms to the requirements<br> defined by the health check.|
|`UNHEALTHY`|The instance is reachable, but does not conform to the requirements<br> defined by the health check.|
|`DRAINING`|The instance is being drained. The existing connections to the instance<br> have time to complete, but the new ones are being refused.|
|`TIMEOUT`|The instance is unreachable i.e. a connection to the application health<br> checking endpoint cannot be established, or the server does not respond<br> within the specified timeout.|

<h2 id="LivenessState">Enum <code>LivenessState</code></h2>

Liveness health check status for Flex instances.

| Options | |
|---|---|
|`LIVENESS_STATE_UNSPECIFIED`|There is no liveness health check for the instance. Only applicable for<br> instances in App Engine standard environment.|
|`UNKNOWN`|The health checking system is aware of the instance but its health is<br> not known at the moment.|
|`HEALTHY`|The instance is reachable i.e. a connection to the application health<br> checking endpoint can be established, and conforms to the requirements<br> defined by the health check.|
|`UNHEALTHY`|The instance is reachable, but does not conform to the requirements<br> defined by the health check.|
|`DRAINING`|The instance is being drained. The existing connections to the instance<br> have time to complete, but the new ones are being refused.|
|`TIMEOUT`|The instance is unreachable i.e. a connection to the application health<br> checking endpoint cannot be established, or the server does not respond<br> within the specified timeout.|

<h2 id="TrafficSplit.ShardBy">Enum <code>TrafficSplit.ShardBy</code></h2>

Available sharding mechanisms.

| Options | |
|---|---|
|`UNSPECIFIED`|Diversion method unspecified.|
|`COOKIE`|Diversion based on a specially named cookie, "GOOGAPPUID." The cookie<br> must be set by the application itself or no diversion will occur.|
|`IP`|Diversion based on applying the modulus operation to a fingerprint<br> of the IP address.|
|`RANDOM`|Diversion based on weighted random assignment. An incoming request is<br> randomly routed to a version in the traffic split, with probability<br> proportional to the version's traffic share.|

<h2 id="ShardBy">Enum <code>ShardBy</code></h2>

Available sharding mechanisms.

| Options | |
|---|---|
|`UNSPECIFIED`|Diversion method unspecified.|
|`COOKIE`|Diversion based on a specially named cookie, "GOOGAPPUID." The cookie<br> must be set by the application itself or no diversion will occur.|
|`IP`|Diversion based on applying the modulus operation to a fingerprint<br> of the IP address.|
|`RANDOM`|Diversion based on weighted random assignment. An incoming request is<br> randomly routed to a version in the traffic split, with probability<br> proportional to the version's traffic share.|

<h2 id="Instance.Liveness.LivenessState">Enum <code>Instance.Liveness.LivenessState</code></h2>

Liveness health check status for Flex instances.

| Options | |
|---|---|
|`LIVENESS_STATE_UNSPECIFIED`|There is no liveness health check for the instance. Only applicable for<br> instances in App Engine standard environment.|
|`UNKNOWN`|The health checking system is aware of the instance but its health is<br> not known at the moment.|
|`HEALTHY`|The instance is reachable i.e. a connection to the application health<br> checking endpoint can be established, and conforms to the requirements<br> defined by the health check.|
|`UNHEALTHY`|The instance is reachable, but does not conform to the requirements<br> defined by the health check.|
|`DRAINING`|The instance is being drained. The existing connections to the instance<br> have time to complete, but the new ones are being refused.|
|`TIMEOUT`|The instance is unreachable i.e. a connection to the application health<br> checking endpoint cannot be established, or the server does not respond<br> within the specified timeout.|

<h2 id="NetworkSettings.IngressTrafficAllowed">Enum <code>NetworkSettings.IngressTrafficAllowed</code></h2>

If unspecified, INGRESS_TRAFFIC_ALLOWED_ALL will be used.

| Options | |
|---|---|
|`INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED`|Unspecified|
|`INGRESS_TRAFFIC_ALLOWED_ALL`|Allow HTTP traffic from public and private sources.|
|`INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY`|Allow HTTP traffic from only private VPC sources.|
|`INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB`|Allow HTTP traffic from private VPC sources and through load balancers.|

<h2 id="IngressTrafficAllowed">Enum <code>IngressTrafficAllowed</code></h2>

If unspecified, INGRESS_TRAFFIC_ALLOWED_ALL will be used.

| Options | |
|---|---|
|`INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED`|Unspecified|
|`INGRESS_TRAFFIC_ALLOWED_ALL`|Allow HTTP traffic from public and private sources.|
|`INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY`|Allow HTTP traffic from only private VPC sources.|
|`INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB`|Allow HTTP traffic from private VPC sources and through load balancers.|



<h2 id="Service">Message <code>Service</code></h2>

A Service resource is a logical component of an application that can share
 state and communicate in a secure fashion with other services.
 For example, an application that handles customer requests might
 include separate services to handle tasks such as backend data
 analysis or API requests from mobile devices. Each service has a
 collection of versions that define a specific set of code used to
 implement the functionality of that service.

| Fields | |
|---|---|
|`name`|**string**<br><br>Full path to the Service resource in the API.<br> Example: `apps/myapp/services/default`.<br><br> @OutputOnly|
|`id`|**string**<br><br>Relative name of the service within the application.<br> Example: `default`.<br><br> @OutputOnly|
|`split`|**[google.appengine.v1beta.TrafficSplit](./index.md#TrafficSplit)**<br><br>Mapping that defines fractional HTTP traffic diversion to<br> different versions within the service.|
|`network_settings`|**[google.appengine.v1beta.NetworkSettings](./index.md#NetworkSettings)**<br><br>Ingress settings for this service. Will apply to all versions.|

<h2 id="TrafficSplit.AllocationsEntry">Message <code>TrafficSplit.AllocationsEntry</code></h2>



| Fields | |
|---|---|
|`key`|**string**<br><br>|
|`value`|**double**<br><br>|

<h2 id="ListServicesRequest">Message <code>ListServicesRequest</code></h2>

Request message for `Services.ListServices`.

| Fields | |
|---|---|
|`parent`|**string**<br><br>Name of the parent Application resource. Example: `apps/myapp`.|
|`page_size`|**int32**<br><br>Maximum results to return per page.|
|`page_token`|**string**<br><br>Continuation token for fetching the next page of results.|

<h2 id="ListServicesResponse">Message <code>ListServicesResponse</code></h2>

Response message for `Services.ListServices`.

| Fields | |
|---|---|
|`services`|**[[]google.appengine.v1beta.Service](./index.md#Service)**<br><br>The services belonging to the requested application.|
|`next_page_token`|**string**<br><br>Continuation token for fetching the next page of results.|

<h2 id="GetServiceRequest">Message <code>GetServiceRequest</code></h2>

Request message for `Services.GetService`.

| Fields | |
|---|---|
|`name`|**string**<br><br>Name of the resource requested. Example: `apps/myapp/services/default`.|

<h2 id="DeleteServiceRequest">Message <code>DeleteServiceRequest</code></h2>

Request message for `Services.DeleteService`.

| Fields | |
|---|---|
|`name`|**string**<br><br>Name of the resource requested. Example: `apps/myapp/services/default`.|

<h2 id="ListInstancesRequest">Message <code>ListInstancesRequest</code></h2>

Request message for `Instances.ListInstances`.

| Fields | |
|---|---|
|`parent`|**string**<br><br>Name of the parent Version resource. Example:<br> `apps/myapp/services/default/versions/v1`.|
|`page_size`|**int32**<br><br>Maximum results to return per page.|
|`page_token`|**string**<br><br>Continuation token for fetching the next page of results.|

<h2 id="ListInstancesResponse">Message <code>ListInstancesResponse</code></h2>

Response message for `Instances.ListInstances`.

| Fields | |
|---|---|
|`instances`|**[[]google.appengine.v1beta.Instance](./index.md#Instance)**<br><br>The instances belonging to the requested version.|
|`next_page_token`|**string**<br><br>Continuation token for fetching the next page of results.|

<h2 id="Instance">Message <code>Instance</code></h2>

An Instance resource is the computing unit that App Engine uses to
 automatically scale an application.

| Fields | |
|---|---|
|`name`|**string**<br><br>Output only. Full path to the Instance resource in the API.<br> Example: `apps/myapp/services/default/versions/v1/instances/instance-1`.|
|`id`|**string**<br><br>Output only. Relative name of the instance within the version.<br> Example: `instance-1`.|
|`app_engine_release`|**string**<br><br>Output only. App Engine release this instance is running on.|
|`availability`|**google.appengine.v1beta.Instance.Availability**<br><br>Output only. Availability of the instance.|
|`vm_name`|**string**<br><br>Output only. Name of the virtual machine where this instance lives. Only applicable<br> for instances in App Engine flexible environment.|
|`vm_zone_name`|**string**<br><br>Output only. Zone where the virtual machine is located. Only applicable for instances<br> in App Engine flexible environment.|
|`vm_id`|**string**<br><br>Output only. Virtual machine ID of this instance. Only applicable for instances in<br> App Engine flexible environment.|
|`start_time`|**[google.protobuf.Timestamp](https://protobuf.dev/reference/protobuf/google.protobuf/#timestamp)**<br><br>Output only. Time that this instance was started.<br><br> @OutputOnly|
|`requests`|**int32**<br><br>Output only. Number of requests since this instance was started.|
|`errors`|**int32**<br><br>Output only. Number of errors since this instance was started.|
|`qps`|**float**<br><br>Output only. Average queries per second (QPS) over the last minute.|
|`average_latency`|**int32**<br><br>Output only. Average latency (ms) over the last minute.|
|`memory_usage`|**int64**<br><br>Output only. Total memory in use (bytes).|
|`vm_status`|**string**<br><br>Output only. Status of the virtual machine where this instance lives. Only applicable<br> for instances in App Engine flexible environment.|
|`vm_debug_enabled`|**bool**<br><br>Output only. Whether this instance is in debug mode. Only applicable for instances in<br> App Engine flexible environment.|
|`vm_ip`|**string**<br><br>Output only. The IP address of this instance. Only applicable for instances in App<br> Engine flexible environment.|
|`vm_liveness`|**google.appengine.v1beta.Instance.Liveness.LivenessState**<br><br>Output only. The liveness health check of this instance. Only applicable for instances<br> in App Engine flexible environment.|

<h2 id="Instance.Liveness">Message <code>Instance.Liveness</code></h2>

Wrapper for LivenessState enum.

| Fields | |
|---|---|

<h2 id="Liveness">Message <code>Liveness</code></h2>

Wrapper for LivenessState enum.

| Fields | |
|---|---|

<h2 id="GetInstanceRequest">Message <code>GetInstanceRequest</code></h2>

Request message for `Instances.GetInstance`.

| Fields | |
|---|---|
|`name`|**string**<br><br>Name of the resource requested. Example:<br> `apps/myapp/services/default/versions/v1/instances/instance-1`.|

<h2 id="NetworkSettings">Message <code>NetworkSettings</code></h2>

A NetworkSettings resource is a container for ingress settings for a version
 or service.

| Fields | |
|---|---|
|`ingress_traffic_allowed`|**google.appengine.v1beta.NetworkSettings.IngressTrafficAllowed**<br><br>The ingress settings for version or service.|

<h2 id="ListIngressRulesRequest">Message <code>ListIngressRulesRequest</code></h2>

Request message for `Firewall.ListIngressRules`.

| Fields | |
|---|---|
|`parent`|**string**<br><br>Name of the Firewall collection to retrieve.<br> Example: `apps/myapp/firewall/ingressRules`.|
|`page_size`|**int32**<br><br>Maximum results to return per page.|
|`page_token`|**string**<br><br>Continuation token for fetching the next page of results.|
|`matching_address`|**string**<br><br>A valid IP Address. If set, only rules matching this address will be<br> returned. The first returned rule will be the rule that fires on requests<br> from this IP.|

<h2 id="TrafficSplit">Message <code>TrafficSplit</code></h2>

Traffic routing configuration for versions within a single service. Traffic
 splits define how traffic directed to the service is assigned to versions.

| Fields | |
|---|---|
|`shard_by`|**google.appengine.v1beta.TrafficSplit.ShardBy**<br><br>Mechanism used to determine which version a request is sent to.<br> The traffic selection algorithm will<br> be stable for either type until allocations are changed.|
|`allocations`|**[[]google.appengine.v1beta.TrafficSplit.AllocationsEntry](./index.md#AllocationsEntry)**<br><br>Mapping from version IDs within the service to fractional<br> (0.000, 1] allocations of traffic for that version. Each version can<br> be specified only once, but some versions in the service may not<br> have any traffic allocation. Services that have traffic allocated<br> cannot be deleted until either the service is deleted or<br> their traffic allocation is removed. Allocations must sum to 1.<br> Up to two decimal place precision is supported for IP-based splits and<br> up to three decimal places is supported for cookie-based splits.|

<h2 id="AllocationsEntry">Message <code>AllocationsEntry</code></h2>



| Fields | |
|---|---|
|`key`|**string**<br><br>|
|`value`|**double**<br><br>|

<h2 id="UpdateServiceRequest">Message <code>UpdateServiceRequest</code></h2>

Request message for `Services.UpdateService`.

| Fields | |
|---|---|
|`name`|**string**<br><br>Name of the resource to update. Example: `apps/myapp/services/default`.|
|`service`|**[google.appengine.v1beta.Service](./index.md#Service)**<br><br>A Service resource containing the updated service. Only fields set in the<br> field mask will be updated.|
|`update_mask`|**[google.protobuf.FieldMask](https://protobuf.dev/reference/protobuf/google.protobuf/#field-mask)**<br><br>Standard field mask for the set of fields to be updated.|
|`migrate_traffic`|**bool**<br><br>Set to `true` to gradually shift traffic to one or more versions that you<br> specify. By default, traffic is shifted immediately.<br> For gradual traffic migration, the target versions<br> must be located within instances that are configured for both<br> [warmup requests](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1beta/apps.services.versions#InboundServiceType)<br> and<br> [automatic scaling](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1beta/apps.services.versions#AutomaticScaling).<br> You must specify the<br> [`shardBy`](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1beta/apps.services#ShardBy)<br> field in the Service resource. Gradual traffic migration is not<br> supported in the App Engine flexible environment. For examples, see<br> [Migrating and Splitting Traffic](https://cloud.google.com/appengine/docs/admin-api/migrating-splitting-traffic).|

<h2 id="DeleteInstanceRequest">Message <code>DeleteInstanceRequest</code></h2>

Request message for `Instances.DeleteInstance`.

| Fields | |
|---|---|
|`name`|**string**<br><br>Name of the resource requested. Example:<br> `apps/myapp/services/default/versions/v1/instances/instance-1`.|

<h2 id="DebugInstanceRequest">Message <code>DebugInstanceRequest</code></h2>

Request message for `Instances.DebugInstance`.

| Fields | |
|---|---|
|`name`|**string**<br><br>Name of the resource requested. Example:<br> `apps/myapp/services/default/versions/v1/instances/instance-1`.|
|`ssh_key`|**string**<br><br>Public SSH key to add to the instance. Examples:<br><br> * `[USERNAME]:ssh-rsa [KEY_VALUE] [USERNAME]`<br> * `[USERNAME]:ssh-rsa [KEY_VALUE] google-ssh {"userName":"[USERNAME]","expireOn":"[EXPIRE_TIME]"}`<br><br> For more information, see<br> [Adding and Removing SSH Keys](https://cloud.google.com/compute/docs/instances/adding-removing-ssh-keys).|

