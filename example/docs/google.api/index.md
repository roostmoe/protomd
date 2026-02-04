# Package google.api

## Index

- [`FieldBehavior`](#E_FieldBehavior) (enum)
- [`LaunchStage`](#E_LaunchStage) (enum)
- [`ClientLibraryOrganization`](#E_ClientLibraryOrganization) (enum)
- [`ClientLibraryDestination`](#E_ClientLibraryDestination) (enum)
- [`ResourceDescriptor`](#M_ResourceDescriptor) (message)
- [`ResourceReference`](#M_ResourceReference) (message)
- [`Http`](#M_Http) (message)
- [`HttpRule`](#M_HttpRule) (message)
- [`CustomHttpPattern`](#M_CustomHttpPattern) (message)
- [`CommonLanguageSettings`](#M_CommonLanguageSettings) (message)
- [`ClientLibrarySettings`](#M_ClientLibrarySettings) (message)
- [`Publishing`](#M_Publishing) (message)
- [`JavaSettings`](#M_JavaSettings) (message)
- [`CppSettings`](#M_CppSettings) (message)
- [`PhpSettings`](#M_PhpSettings) (message)
- [`PythonSettings`](#M_PythonSettings) (message)
- [`NodeSettings`](#M_NodeSettings) (message)
- [`DotnetSettings`](#M_DotnetSettings) (message)
- [`RubySettings`](#M_RubySettings) (message)
- [`GoSettings`](#M_GoSettings) (message)
- [`MethodSettings`](#M_MethodSettings) (message)
- [`SelectiveGapicGeneration`](#M_SelectiveGapicGeneration) (message)

---

<h2 id="E_FieldBehavior">Enum <code>FieldBehavior</code></h2>



| Options | |
|---|---|
|`FIELD_BEHAVIOR_UNSPECIFIED`||
|`OPTIONAL`||
|`REQUIRED`||
|`OUTPUT_ONLY`||
|`INPUT_ONLY`||
|`IMMUTABLE`||
|`UNORDERED_LIST`||
|`NON_EMPTY_DEFAULT`||
|`IDENTIFIER`||

<h2 id="E_LaunchStage">Enum <code>LaunchStage</code></h2>



| Options | |
|---|---|
|`LAUNCH_STAGE_UNSPECIFIED`||
|`UNIMPLEMENTED`||
|`PRELAUNCH`||
|`EARLY_ACCESS`||
|`ALPHA`||
|`BETA`||
|`GA`||
|`DEPRECATED`||

<h2 id="E_ClientLibraryOrganization">Enum <code>ClientLibraryOrganization</code></h2>

Required information for every language.

| Options | |
|---|---|
|`CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED`|Link to automatically generated reference documentation.  Example:<br/>https://cloud.google.com/nodejs/docs/reference/asset/latest|
|`CLOUD`|The destination where API teams want this client library to be published.|
|`ADS`|Configuration for which RPCs should be generated in the GAPIC client.|
|`PHOTOS`||
|`STREET_VIEW`||
|`SHOPPING`||
|`GEO`||
|`GENERATIVE_AI`||

<h2 id="E_ClientLibraryDestination">Enum <code>ClientLibraryDestination</code></h2>

Details about how and where to publish client libraries.

| Options | |
|---|---|
|`CLIENT_LIBRARY_DESTINATION_UNSPECIFIED`|Version of the API to apply these settings to. This is the full protobuf<br/>package for the API, ending in the version element.<br/>Examples: "google.cloud.speech.v1" and "google.spanner.admin.database.v1".|
|`GITHUB`|Launch stage of this version of the API.|
|`PACKAGE_MANAGER`|When using transport=rest, the client request will encode enums as<br/>numbers rather than strings.|

<h2 id="M_ResourceDescriptor">Message <code>ResourceDescriptor</code></h2>

A simple descriptor of a resource type.

 ResourceDescriptor annotates a resource message (either by means of a
 protobuf annotation or use in the service config), and associates the
 resource's schema, the resource type, and the pattern of the resource name.

 Example:

     message Topic {
       // Indicates this message defines a resource schema.
       // Declares the resource type in the format of {service}/{kind}.
       // For Kubernetes resources, the format is {api group}/{kind}.
       option (google.api.resource) = {
         type: "pubsub.googleapis.com/Topic"
         pattern: "projects/{project}/topics/{topic}"
       };
     }

 The ResourceDescriptor Yaml config will look like:

     resources:
     - type: "pubsub.googleapis.com/Topic"
       pattern: "projects/{project}/topics/{topic}"

 Sometimes, resources have multiple patterns, typically because they can
 live under multiple parents.

 Example:

     message LogEntry {
       option (google.api.resource) = {
         type: "logging.googleapis.com/LogEntry"
         pattern: "projects/{project}/logs/{log}"
         pattern: "folders/{folder}/logs/{log}"
         pattern: "organizations/{organization}/logs/{log}"
         pattern: "billingAccounts/{billing_account}/logs/{log}"
       };
     }

 The ResourceDescriptor Yaml config will look like:

     resources:
     - type: 'logging.googleapis.com/LogEntry'
       pattern: "projects/{project}/logs/{log}"
       pattern: "folders/{folder}/logs/{log}"
       pattern: "organizations/{organization}/logs/{log}"
       pattern: "billingAccounts/{billing_account}/logs/{log}"

| Fields | |
|---|---|
|`type`|The resource type. It must be in the format of<br/>{service_name}/{resource_type_kind}. The `resource_type_kind` must be<br/>singular and must not include version numbers.<br/><br/>Example: `storage.googleapis.com/Bucket`<br/><br/>The value of the resource_type_kind must follow the regular expression<br/>/[A-Za-z][a-zA-Z0-9]+/. It should start with an upper case character and<br/>should use PascalCase (UpperCamelCase). The maximum number of<br/>characters allowed for the `resource_type_kind` is 100.|
|`pattern`|Optional. The relative resource name pattern associated with this resource<br/>type. The DNS prefix of the full resource name shouldn't be specified here.<br/><br/>The path pattern must follow the syntax, which aligns with HTTP binding<br/>syntax:<br/><br/>Template = Segment { "/" Segment } ;<br/>Segment = LITERAL | Variable ;<br/>Variable = "{" LITERAL "}" ;<br/><br/>Examples:<br/><br/>- "projects/{project}/topics/{topic}"<br/>- "projects/{project}/knowledgeBases/{knowledge_base}"<br/><br/>The components in braces correspond to the IDs for each resource in the<br/>hierarchy. It is expected that, if multiple patterns are provided,<br/>the same component name (e.g. "project") refers to IDs of the same<br/>type of resource.|
|`name_field`|Optional. The field on the resource that designates the resource name<br/>field. If omitted, this is assumed to be "name".|
|`history`|Optional. The historical or future-looking state of the resource pattern.<br/><br/>Example:<br/><br/>// The InspectTemplate message originally only supported resource<br/>// names with organization, and project was added later.<br/>message InspectTemplate {<br/>option (google.api.resource) = {<br/>type: "dlp.googleapis.com/InspectTemplate"<br/>pattern:<br/>"organizations/{organization}/inspectTemplates/{inspect_template}"<br/>pattern: "projects/{project}/inspectTemplates/{inspect_template}"<br/>history: ORIGINALLY_SINGLE_PATTERN<br/>};<br/>}|
|`plural`|The plural name used in the resource name and permission names, such as<br/>'projects' for the resource name of 'projects/{project}' and the permission<br/>name of 'cloudresourcemanager.googleapis.com/projects.get'. One exception<br/>to this is for Nested Collections that have stuttering names, as defined<br/>in [AIP-122](https://google.aip.dev/122#nested-collections), where the<br/>collection ID in the resource name pattern does not necessarily directly<br/>match the `plural` value.<br/><br/>It is the same concept of the `plural` field in k8s CRD spec<br/>https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/<br/><br/>Note: The plural form is required even for singleton resources. See<br/>https://aip.dev/156|
|`singular`|The same concept of the `singular` field in k8s CRD spec<br/>https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/<br/>Such as "project" for the `resourcemanager.googleapis.com/Project` type.|
|`style`|Style flag(s) for this resource.<br/>These indicate that a resource is expected to conform to a given<br/>style. See the specific style flags for additional information.|

<h2 id="M_ResourceReference">Message <code>ResourceReference</code></h2>

Defines a proto annotation that describes a string field that refers to
 an API resource.

| Fields | |
|---|---|
|`type`|The resource type that the annotated field references.<br/><br/>Example:<br/><br/>message Subscription {<br/>string topic = 2 [(google.api.resource_reference) = {<br/>type: "pubsub.googleapis.com/Topic"<br/>}];<br/>}<br/><br/>Occasionally, a field may reference an arbitrary resource. In this case,<br/>APIs use the special value * in their resource reference.<br/><br/>Example:<br/><br/>message GetIamPolicyRequest {<br/>string resource = 2 [(google.api.resource_reference) = {<br/>type: "*"<br/>}];<br/>}|
|`child_type`|The resource type of a child collection that the annotated field<br/>references. This is useful for annotating the `parent` field that<br/>doesn't have a fixed resource type.<br/><br/>Example:<br/><br/>message ListLogEntriesRequest {<br/>string parent = 1 [(google.api.resource_reference) = {<br/>child_type: "logging.googleapis.com/LogEntry"<br/>};<br/>}|

<h2 id="M_Http">Message <code>Http</code></h2>

Defines the HTTP configuration for an API service. It contains a list of
 [HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method
 to one or more HTTP REST API methods.

| Fields | |
|---|---|
|`rules`|A list of HTTP configuration rules that apply to individual API methods.<br/><br/>**NOTE:** All service configuration rules follow "last one wins" order.|
|`fully_decode_reserved_expansion`|When set to true, URL path parameters will be fully URI-decoded except in<br/>cases of single segment matches in reserved expansion, where "%2F" will be<br/>left encoded.<br/><br/>The default behavior is to not decode RFC 6570 reserved characters in multi<br/>segment matches.|

<h2 id="M_HttpRule">Message <code>HttpRule</code></h2>

gRPC Transcoding

 gRPC Transcoding is a feature for mapping between a gRPC method and one or
 more HTTP REST endpoints. It allows developers to build a single API service
 that supports both gRPC APIs and REST APIs. Many systems, including [Google
 APIs](https://github.com/googleapis/googleapis),
 [Cloud Endpoints](https://cloud.google.com/endpoints), [gRPC
 Gateway](https://github.com/grpc-ecosystem/grpc-gateway),
 and [Envoy](https://github.com/envoyproxy/envoy) proxy support this feature
 and use it for large scale production services.

 `HttpRule` defines the schema of the gRPC/REST mapping. The mapping specifies
 how different portions of the gRPC request message are mapped to the URL
 path, URL query parameters, and HTTP request body. It also controls how the
 gRPC response message is mapped to the HTTP response body. `HttpRule` is
 typically specified as an `google.api.http` annotation on the gRPC method.

 Each mapping specifies a URL path template and an HTTP method. The path
 template may refer to one or more fields in the gRPC request message, as long
 as each field is a non-repeated field with a primitive (non-message) type.
 The path template controls how fields of the request message are mapped to
 the URL path.

 Example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
             get: "/v1/{name=messages/*}"
         };
       }
     }
     message GetMessageRequest {
       string name = 1; // Mapped to URL path.
     }
     message Message {
       string text = 1; // The resource content.
     }

 This enables an HTTP REST to gRPC mapping as below:

 - HTTP: `GET /v1/messages/123456`
 - gRPC: `GetMessage(name: "messages/123456")`

 Any fields in the request message which are not bound by the path template
 automatically become HTTP query parameters if there is no HTTP request body.
 For example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
             get:"/v1/messages/{message_id}"
         };
       }
     }
     message GetMessageRequest {
       message SubMessage {
         string subfield = 1;
       }
       string message_id = 1; // Mapped to URL path.
       int64 revision = 2;    // Mapped to URL query parameter `revision`.
       SubMessage sub = 3;    // Mapped to URL query parameter `sub.subfield`.
     }

 This enables a HTTP JSON to RPC mapping as below:

 - HTTP: `GET /v1/messages/123456?revision=2&sub.subfield=foo`
 - gRPC: `GetMessage(message_id: "123456" revision: 2 sub:
 SubMessage(subfield: "foo"))`

 Note that fields which are mapped to URL query parameters must have a
 primitive type or a repeated primitive type or a non-repeated message type.
 In the case of a repeated type, the parameter can be repeated in the URL
 as `...?param=A&param=B`. In the case of a message type, each field of the
 message is mapped to a separate parameter, such as
 `...?foo.a=A&foo.b=B&foo.c=C`.

 For HTTP methods that allow a request body, the `body` field
 specifies the mapping. Consider a REST update method on the
 message resource collection:

     service Messaging {
       rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
         option (google.api.http) = {
           patch: "/v1/messages/{message_id}"
           body: "message"
         };
       }
     }
     message UpdateMessageRequest {
       string message_id = 1; // mapped to the URL
       Message message = 2;   // mapped to the body
     }

 The following HTTP JSON to RPC mapping is enabled, where the
 representation of the JSON in the request body is determined by
 protos JSON encoding:

 - HTTP: `PATCH /v1/messages/123456 { "text": "Hi!" }`
 - gRPC: `UpdateMessage(message_id: "123456" message { text: "Hi!" })`

 The special name `*` can be used in the body mapping to define that
 every field not bound by the path template should be mapped to the
 request body.  This enables the following alternative definition of
 the update method:

     service Messaging {
       rpc UpdateMessage(Message) returns (Message) {
         option (google.api.http) = {
           patch: "/v1/messages/{message_id}"
           body: "*"
         };
       }
     }
     message Message {
       string message_id = 1;
       string text = 2;
     }


 The following HTTP JSON to RPC mapping is enabled:

 - HTTP: `PATCH /v1/messages/123456 { "text": "Hi!" }`
 - gRPC: `UpdateMessage(message_id: "123456" text: "Hi!")`

 Note that when using `*` in the body mapping, it is not possible to
 have HTTP parameters, as all fields not bound by the path end in
 the body. This makes this option more rarely used in practice when
 defining REST APIs. The common usage of `*` is in custom methods
 which don't use the URL at all for transferring data.

 It is possible to define multiple HTTP methods for one RPC by using
 the `additional_bindings` option. Example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
           get: "/v1/messages/{message_id}"
           additional_bindings {
             get: "/v1/users/{user_id}/messages/{message_id}"
           }
         };
       }
     }
     message GetMessageRequest {
       string message_id = 1;
       string user_id = 2;
     }

 This enables the following two alternative HTTP JSON to RPC mappings:

 - HTTP: `GET /v1/messages/123456`
 - gRPC: `GetMessage(message_id: "123456")`

 - HTTP: `GET /v1/users/me/messages/123456`
 - gRPC: `GetMessage(user_id: "me" message_id: "123456")`

 Rules for HTTP mapping

 1. Leaf request fields (recursive expansion nested messages in the request
    message) are classified into three categories:
    - Fields referred by the path template. They are passed via the URL path.
    - Fields referred by the [HttpRule.body][google.api.HttpRule.body]. They
    are passed via the HTTP
      request body.
    - All other fields are passed via the URL query parameters, and the
      parameter name is the field path in the request message. A repeated
      field can be represented as multiple query parameters under the same
      name.
  2. If [HttpRule.body][google.api.HttpRule.body] is "*", there is no URL
  query parameter, all fields
     are passed via URL path and HTTP request body.
  3. If [HttpRule.body][google.api.HttpRule.body] is omitted, there is no HTTP
  request body, all
     fields are passed via URL path and URL query parameters.

 Path template syntax

     Template = "/" Segments [ Verb ] ;
     Segments = Segment { "/" Segment } ;
     Segment  = "*" | "**" | LITERAL | Variable ;
     Variable = "{" FieldPath [ "=" Segments ] "}" ;
     FieldPath = IDENT { "." IDENT } ;
     Verb     = ":" LITERAL ;

 The syntax `*` matches a single URL path segment. The syntax `**` matches
 zero or more URL path segments, which must be the last part of the URL path
 except the `Verb`.

 The syntax `Variable` matches part of the URL path as specified by its
 template. A variable template must not contain other variables. If a variable
 matches a single path segment, its template may be omitted, e.g. `{var}`
 is equivalent to `{var=*}`.

 The syntax `LITERAL` matches literal text in the URL path. If the `LITERAL`
 contains any reserved character, such characters should be percent-encoded
 before the matching.

 If a variable contains exactly one path segment, such as `"{var}"` or
 `"{var=*}"`, when such a variable is expanded into a URL path on the client
 side, all characters except `[-_.~0-9a-zA-Z]` are percent-encoded. The
 server side does the reverse decoding. Such variables show up in the
 [Discovery
 Document](https://developers.google.com/discovery/v1/reference/apis) as
 `{var}`.

 If a variable contains multiple path segments, such as `"{var=foo/*}"`
 or `"{var=**}"`, when such a variable is expanded into a URL path on the
 client side, all characters except `[-_.~/0-9a-zA-Z]` are percent-encoded.
 The server side does the reverse decoding, except "%2F" and "%2f" are left
 unchanged. Such variables show up in the
 [Discovery
 Document](https://developers.google.com/discovery/v1/reference/apis) as
 `{+var}`.

 Using gRPC API Service Configuration

 gRPC API Service Configuration (service config) is a configuration language
 for configuring a gRPC service to become a user-facing product. The
 service config is simply the YAML representation of the `google.api.Service`
 proto message.

 As an alternative to annotating your proto file, you can configure gRPC
 transcoding in your service config YAML files. You do this by specifying a
 `HttpRule` that maps the gRPC method to a REST endpoint, achieving the same
 effect as the proto annotation. This can be particularly useful if you
 have a proto that is reused in multiple services. Note that any transcoding
 specified in the service config will override any matching transcoding
 configuration in the proto.

 The following example selects a gRPC method and applies an `HttpRule` to it:

     http:
       rules:
         - selector: example.v1.Messaging.GetMessage
           get: /v1/messages/{message_id}/{sub.subfield}

 Special notes

 When gRPC Transcoding is used to map a gRPC to JSON REST endpoints, the
 proto to JSON conversion must follow the [proto3
 specification](https://developers.google.com/protocol-buffers/docs/proto3#json).

 While the single segment variable follows the semantics of
 [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2 Simple String
 Expansion, the multi segment variable **does not** follow RFC 6570 Section
 3.2.3 Reserved Expansion. The reason is that the Reserved Expansion
 does not expand special characters like `?` and `#`, which would lead
 to invalid URLs. As the result, gRPC Transcoding uses a custom encoding
 for multi segment variables.

 The path variables **must not** refer to any repeated or mapped field,
 because client libraries are not capable of handling such variable expansion.

 The path variables **must not** capture the leading "/" character. The reason
 is that the most common use case "{var}" does not capture the leading "/"
 character. For consistency, all path variables must share the same behavior.

 Repeated message fields must not be mapped to URL query parameters, because
 no client library can support such complicated mapping.

 If an API needs to use a JSON array for request or response body, it can map
 the request or response body to a repeated field. However, some gRPC
 Transcoding implementations may not support this feature.

| Fields | |
|---|---|
|`selector`|Selects a method to which this rule applies.<br/><br/>Refer to [selector][google.api.DocumentationRule.selector] for syntax<br/>details.|
|`get`|Maps to HTTP GET. Used for listing and getting information about<br/>resources.|
|`put`|Maps to HTTP PUT. Used for replacing a resource.|
|`post`|Maps to HTTP POST. Used for creating a resource or performing an action.|
|`delete`|Maps to HTTP DELETE. Used for deleting a resource.|
|`patch`|Maps to HTTP PATCH. Used for updating a resource.|
|`custom`|The custom pattern is used for specifying an HTTP method that is not<br/>included in the `pattern` field, such as HEAD, or "*" to leave the<br/>HTTP method unspecified for this rule. The wild-card rule is useful<br/>for services that provide content to Web (HTML) clients.|
|`body`|The name of the request field whose value is mapped to the HTTP request<br/>body, or `*` for mapping all request fields not captured by the path<br/>pattern to the HTTP body, or omitted for not having any HTTP request body.<br/><br/>NOTE: the referred field must be present at the top-level of the request<br/>message type.|
|`response_body`|Optional. The name of the response field whose value is mapped to the HTTP<br/>response body. When omitted, the entire response message will be used<br/>as the HTTP response body.<br/><br/>NOTE: The referred field must be present at the top-level of the response<br/>message type.|
|`additional_bindings`|Additional HTTP bindings for the selector. Nested bindings must<br/>not contain an `additional_bindings` field themselves (that is,<br/>the nesting may only be one level deep).|

<h2 id="M_CustomHttpPattern">Message <code>CustomHttpPattern</code></h2>

A custom pattern is used for defining custom HTTP verb.

| Fields | |
|---|---|
|`kind`|The name of this custom HTTP verb.|
|`path`|The path matched by this custom verb.|

<h2 id="M_CommonLanguageSettings">Message <code>CommonLanguageSettings</code></h2>

Required information for every language.

| Fields | |
|---|---|
|`reference_docs_uri`|Link to automatically generated reference documentation.  Example:<br/>https://cloud.google.com/nodejs/docs/reference/asset/latest|
|`destinations`|The destination where API teams want this client library to be published.|
|`selective_gapic_generation`|Configuration for which RPCs should be generated in the GAPIC client.|

<h2 id="M_ClientLibrarySettings">Message <code>ClientLibrarySettings</code></h2>

Details about how and where to publish client libraries.

| Fields | |
|---|---|
|`version`|Version of the API to apply these settings to. This is the full protobuf<br/>package for the API, ending in the version element.<br/>Examples: "google.cloud.speech.v1" and "google.spanner.admin.database.v1".|
|`launch_stage`|Launch stage of this version of the API.|
|`rest_numeric_enums`|When using transport=rest, the client request will encode enums as<br/>numbers rather than strings.|
|`java_settings`|Settings for legacy Java features, supported in the Service YAML.|
|`cpp_settings`|Settings for C++ client libraries.|
|`php_settings`|Settings for PHP client libraries.|
|`python_settings`|Settings for Python client libraries.|
|`node_settings`|Settings for Node client libraries.|
|`dotnet_settings`|Settings for .NET client libraries.|
|`ruby_settings`|Settings for Ruby client libraries.|
|`go_settings`|Settings for Go client libraries.|

<h2 id="M_Publishing">Message <code>Publishing</code></h2>

This message configures the settings for publishing [Google Cloud Client
 libraries](https://cloud.google.com/apis/docs/cloud-client-libraries)
 generated from the service config.

| Fields | |
|---|---|
|`method_settings`|A list of API method settings, e.g. the behavior for methods that use the<br/>long-running operation pattern.|
|`new_issue_uri`|Link to a *public* URI where users can report issues.  Example:<br/>https://issuetracker.google.com/issues/new?component=190865&template=1161103|
|`documentation_uri`|Link to product home page.  Example:<br/>https://cloud.google.com/asset-inventory/docs/overview|
|`api_short_name`|Used as a tracking tag when collecting data about the APIs developer<br/>relations artifacts like docs, packages delivered to package managers,<br/>etc.  Example: "speech".|
|`github_label`|GitHub label to apply to issues and pull requests opened for this API.|
|`codeowner_github_teams`|GitHub teams to be added to CODEOWNERS in the directory in GitHub<br/>containing source code for the client libraries for this API.|
|`doc_tag_prefix`|A prefix used in sample code when demarking regions to be included in<br/>documentation.|
|`organization`|For whom the client library is being published.|
|`library_settings`|Client library settings.  If the same version string appears multiple<br/>times in this list, then the last one wins.  Settings from earlier<br/>settings with the same version string are discarded.|
|`proto_reference_documentation_uri`|Optional link to proto reference documentation.  Example:<br/>https://cloud.google.com/pubsub/lite/docs/reference/rpc|
|`rest_reference_documentation_uri`|Optional link to REST reference documentation.  Example:<br/>https://cloud.google.com/pubsub/lite/docs/reference/rest|

<h2 id="M_JavaSettings">Message <code>JavaSettings</code></h2>

Settings for Java client libraries.

| Fields | |
|---|---|
|`library_package`|The package name to use in Java. Clobbers the java_package option<br/>set in the protobuf. This should be used **only** by APIs<br/>who have already set the language_settings.java.package_name" field<br/>in gapic.yaml. API teams should use the protobuf java_package option<br/>where possible.<br/><br/>Example of a YAML configuration::<br/><br/>publishing:<br/>java_settings:<br/>library_package: com.google.cloud.pubsub.v1|
|`service_class_names`|Configure the Java class name to use instead of the service's for its<br/>corresponding generated GAPIC client. Keys are fully-qualified<br/>service names as they appear in the protobuf (including the full<br/>the language_settings.java.interface_names" field in gapic.yaml. API<br/>teams should otherwise use the service name as it appears in the<br/>protobuf.<br/><br/>Example of a YAML configuration::<br/><br/>publishing:<br/>java_settings:<br/>service_class_names:<br/>- google.pubsub.v1.Publisher: TopicAdmin<br/>- google.pubsub.v1.Subscriber: SubscriptionAdmin|
|`common`|Some settings.|

<h2 id="M_CppSettings">Message <code>CppSettings</code></h2>

Settings for C++ client libraries.

| Fields | |
|---|---|
|`common`|Some settings.|

<h2 id="M_PhpSettings">Message <code>PhpSettings</code></h2>

Settings for Php client libraries.

| Fields | |
|---|---|
|`common`|Some settings.|

<h2 id="M_PythonSettings">Message <code>PythonSettings</code></h2>

Settings for Python client libraries.

| Fields | |
|---|---|
|`common`|Some settings.|
|`experimental_features`|Experimental features to be included during client library generation.|

<h2 id="M_NodeSettings">Message <code>NodeSettings</code></h2>

Settings for Node client libraries.

| Fields | |
|---|---|
|`common`|Some settings.|

<h2 id="M_DotnetSettings">Message <code>DotnetSettings</code></h2>

Settings for Dotnet client libraries.

| Fields | |
|---|---|
|`common`|Some settings.|
|`renamed_services`|Map from original service names to renamed versions.<br/>This is used when the default generated types<br/>would cause a naming conflict. (Neither name is<br/>fully-qualified.)<br/>Example: Subscriber to SubscriberServiceApi.|
|`renamed_resources`|Map from full resource types to the effective short name<br/>for the resource. This is used when otherwise resource<br/>named from different services would cause naming collisions.<br/>Example entry:<br/>"datalabeling.googleapis.com/Dataset": "DataLabelingDataset"|
|`ignored_resources`|List of full resource types to ignore during generation.<br/>This is typically used for API-specific Location resources,<br/>which should be handled by the generator as if they were actually<br/>the common Location resources.<br/>Example entry: "documentai.googleapis.com/Location"|
|`forced_namespace_aliases`|Namespaces which must be aliased in snippets due to<br/>a known (but non-generator-predictable) naming collision|
|`handwritten_signatures`|Method signatures (in the form "service.method(signature)")<br/>which are provided separately, so shouldn't be generated.<br/>Snippets *calling* these methods are still generated, however.|

<h2 id="M_RubySettings">Message <code>RubySettings</code></h2>

Settings for Ruby client libraries.

| Fields | |
|---|---|
|`common`|Some settings.|

<h2 id="M_GoSettings">Message <code>GoSettings</code></h2>

Settings for Go client libraries.

| Fields | |
|---|---|
|`common`|Some settings.|
|`renamed_services`|Map of service names to renamed services. Keys are the package relative<br/>service names and values are the name to be used for the service client<br/>and call options.<br/><br/>publishing:<br/>go_settings:<br/>renamed_services:<br/>Publisher: TopicAdmin|

<h2 id="M_MethodSettings">Message <code>MethodSettings</code></h2>

Describes the generator configuration for a method.

| Fields | |
|---|---|
|`selector`|The fully qualified name of the method, for which the options below apply.<br/>This is used to find the method to apply the options.<br/><br/>Example:<br/><br/>publishing:<br/>method_settings:<br/>- selector: google.storage.control.v2.StorageControl.CreateFolder<br/># method settings for CreateFolder...|
|`long_running`|Describes settings to use for long-running operations when generating<br/>API methods for RPCs. Complements RPCs that use the annotations in<br/>google/longrunning/operations.proto.<br/><br/>Example of a YAML configuration::<br/><br/>publishing:<br/>method_settings:<br/>- selector: google.cloud.speech.v2.Speech.BatchRecognize<br/>long_running:<br/>initial_poll_delay: 60s # 1 minute<br/>poll_delay_multiplier: 1.5<br/>max_poll_delay: 360s # 6 minutes<br/>total_poll_timeout: 54000s # 90 minutes|
|`auto_populated_fields`|List of top-level fields of the request message, that should be<br/>automatically populated by the client libraries based on their<br/>(google.api.field_info).format. Currently supported format: UUID4.<br/><br/>Example of a YAML configuration:<br/><br/>publishing:<br/>method_settings:<br/>- selector: google.example.v1.ExampleService.CreateExample<br/>auto_populated_fields:<br/>- request_id|

<h2 id="M_SelectiveGapicGeneration">Message <code>SelectiveGapicGeneration</code></h2>

This message is used to configure the generation of a subset of the RPCs in
 a service for client libraries.

| Fields | |
|---|---|
|`methods`|An allowlist of the fully qualified names of RPCs that should be included<br/>on public client surfaces.|
|`generate_omitted_as_internal`|Setting this to true indicates to the client generators that methods<br/>that would be excluded from the generation should instead be generated<br/>in a way that indicates these methods should not be consumed by<br/>end users. How this is expressed is up to individual language<br/>implementations to decide. Some examples may be: added annotations,<br/>obfuscated identifiers, or other language idiomatic patterns.|

