# Package google.protobuf

## Index

- [`Edition`](#E_Edition) (enum)
- [`SymbolVisibility`](#E_SymbolVisibility) (enum)
- [`FileDescriptorSet`](#M_FileDescriptorSet) (message)
- [`FileDescriptorProto`](#M_FileDescriptorProto) (message)
- [`DescriptorProto`](#M_DescriptorProto) (message)
- [`ExtensionRangeOptions`](#M_ExtensionRangeOptions) (message)
- [`FieldDescriptorProto`](#M_FieldDescriptorProto) (message)
- [`OneofDescriptorProto`](#M_OneofDescriptorProto) (message)
- [`EnumDescriptorProto`](#M_EnumDescriptorProto) (message)
- [`EnumValueDescriptorProto`](#M_EnumValueDescriptorProto) (message)
- [`ServiceDescriptorProto`](#M_ServiceDescriptorProto) (message)
- [`MethodDescriptorProto`](#M_MethodDescriptorProto) (message)
- [`FileOptions`](#M_FileOptions) (message)
- [`MessageOptions`](#M_MessageOptions) (message)
- [`FieldOptions`](#M_FieldOptions) (message)
- [`OneofOptions`](#M_OneofOptions) (message)
- [`EnumOptions`](#M_EnumOptions) (message)
- [`EnumValueOptions`](#M_EnumValueOptions) (message)
- [`ServiceOptions`](#M_ServiceOptions) (message)
- [`MethodOptions`](#M_MethodOptions) (message)
- [`UninterpretedOption`](#M_UninterpretedOption) (message)
- [`FeatureSet`](#M_FeatureSet) (message)
- [`FeatureSetDefaults`](#M_FeatureSetDefaults) (message)
- [`SourceCodeInfo`](#M_SourceCodeInfo) (message)
- [`GeneratedCodeInfo`](#M_GeneratedCodeInfo) (message)
- [`Timestamp`](#M_Timestamp) (message)
- [`Duration`](#M_Duration) (message)
- [`Any`](#M_Any) (message)
- [`Empty`](#M_Empty) (message)
- [`FieldMask`](#M_FieldMask) (message)

---

<h2 id="E_Edition">Enum <code>Edition</code></h2>

The protocol compiler can output a FileDescriptorSet containing the .proto
 files it parses.

| Options | |
|---|---|
|`EDITION_UNKNOWN`||
|`EDITION_LEGACY`||
|`EDITION_PROTO2`||
|`EDITION_PROTO3`||
|`EDITION_2023`||
|`EDITION_2024`||
|`EDITION_1_TEST_ONLY`||
|`EDITION_2_TEST_ONLY`||
|`EDITION_99997_TEST_ONLY`||
|`EDITION_99998_TEST_ONLY`||
|`EDITION_99999_TEST_ONLY`||
|`EDITION_MAX`||

<h2 id="E_SymbolVisibility">Enum <code>SymbolVisibility</code></h2>

Describes a complete .proto file.

| Options | |
|---|---|
|`VISIBILITY_UNSET`||
|`VISIBILITY_LOCAL`||
|`VISIBILITY_EXPORT`|Names of files imported by this file.|

<h2 id="M_FileDescriptorSet">Message <code>FileDescriptorSet</code></h2>

The protocol compiler can output a FileDescriptorSet containing the .proto
 files it parses.

| Fields | |
|---|---|
|`file`||

<h2 id="M_FileDescriptorProto">Message <code>FileDescriptorProto</code></h2>

Describes a complete .proto file.

| Fields | |
|---|---|
|`name`||
|`package`||
|`dependency`|Names of files imported by this file.|
|`public_dependency`|Indexes of the public imported files in the dependency list above.|
|`weak_dependency`|Indexes of the weak imported files in the dependency list.<br/>For Google-internal migration only. Do not use.|
|`option_dependency`|Names of files imported by this file purely for the purpose of providing<br/>option extensions. These are excluded from the dependency list above.|
|`message_type`|All top-level definitions in this file.|
|`enum_type`||
|`service`||
|`extension`||
|`options`||
|`source_code_info`|This field contains optional information about the original source code.<br/>You may safely remove this entire field without harming runtime<br/>functionality of the descriptors -- the information is needed only by<br/>development tools.|
|`syntax`|The syntax of the proto file.<br/>The supported values are "proto2", "proto3", and "editions".<br/><br/>If `edition` is present, this value must be "editions".<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|
|`edition`|The edition of the proto file.<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|

<h2 id="M_DescriptorProto">Message <code>DescriptorProto</code></h2>

Describes a message type.

| Fields | |
|---|---|
|`name`||
|`field`||
|`extension`||
|`nested_type`||
|`enum_type`||
|`extension_range`||
|`oneof_decl`||
|`options`||
|`reserved_range`||
|`reserved_name`|Reserved field names, which may not be used by fields in the same message.<br/>A given name may only be reserved once.|
|`visibility`|Support for `export` and `local` keywords on enums.|

<h2 id="M_ExtensionRangeOptions">Message <code>ExtensionRangeOptions</code></h2>



| Fields | |
|---|---|
|`uninterpreted_option`|The parser stores options it doesn't recognize here. See above.|
|`declaration`|For external users: DO NOT USE. We are in the process of open sourcing<br/>extension declaration and executing internal cleanups before it can be<br/>used externally.|
|`features`|Any features defined in the specific edition.|
|`verification`|The verification state of the range.<br/>TODO: flip the default to DECLARATION once all empty ranges<br/>are marked as UNVERIFIED.|

<h2 id="M_FieldDescriptorProto">Message <code>FieldDescriptorProto</code></h2>

Describes a field within a message.

| Fields | |
|---|---|
|`name`||
|`number`||
|`label`||
|`type`|If type_name is set, this need not be set.  If both this and type_name<br/>are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP.|
|`type_name`|For message and enum types, this is the name of the type.  If the name<br/>starts with a '.', it is fully-qualified.  Otherwise, C++-like scoping<br/>rules are used to find the type (i.e. first the nested types within this<br/>message are searched, then within the parent, on up to the root<br/>namespace).|
|`extendee`|For extensions, this is the name of the type being extended.  It is<br/>resolved in the same manner as type_name.|
|`default_value`|For numeric types, contains the original text representation of the value.<br/>For booleans, "true" or "false".<br/>For strings, contains the default text contents (not escaped in any way).<br/>For bytes, contains the C escaped value.  All bytes >= 128 are escaped.|
|`oneof_index`|If set, gives the index of a oneof in the containing type's oneof_decl<br/>list.  This field is a member of that oneof.|
|`json_name`|JSON name of this field. The value is set by protocol compiler. If the<br/>user has set a "json_name" option on this field, that option's value<br/>will be used. Otherwise, it's deduced from the field's name by converting<br/>it to camelCase.|
|`options`||
|`proto3_optional`|If true, this is a proto3 "optional". When a proto3 field is optional, it<br/>tracks presence regardless of field type.<br/><br/>When proto3_optional is true, this field must belong to a oneof to signal<br/>to old proto3 clients that presence is tracked for this field. This oneof<br/>is known as a "synthetic" oneof, and this field must be its sole member<br/>(each proto3 optional field gets its own synthetic oneof). Synthetic oneofs<br/>exist in the descriptor only, and do not generate any API. Synthetic oneofs<br/>must be ordered after all "real" oneofs.<br/><br/>For message fields, proto3_optional doesn't create any semantic change,<br/>since non-repeated message fields always track presence. However it still<br/>indicates the semantic detail of whether the user wrote "optional" or not.<br/>This can be useful for round-tripping the .proto file. For consistency we<br/>give message fields a synthetic oneof also, even though it is not required<br/>to track presence. This is especially important because the parser can't<br/>tell if a field is a message or an enum, so it must always create a<br/>synthetic oneof.<br/><br/>Proto2 optional fields do not set this flag, because they already indicate<br/>optional with `LABEL_OPTIONAL`.|

<h2 id="M_OneofDescriptorProto">Message <code>OneofDescriptorProto</code></h2>

Describes a oneof.

| Fields | |
|---|---|
|`name`||
|`options`||

<h2 id="M_EnumDescriptorProto">Message <code>EnumDescriptorProto</code></h2>

Describes an enum type.

| Fields | |
|---|---|
|`name`||
|`value`||
|`options`||
|`reserved_range`|Range of reserved numeric values. Reserved numeric values may not be used<br/>by enum values in the same enum declaration. Reserved ranges may not<br/>overlap.|
|`reserved_name`|Reserved enum value names, which may not be reused. A given name may only<br/>be reserved once.|
|`visibility`|Support for `export` and `local` keywords on enums.|

<h2 id="M_EnumValueDescriptorProto">Message <code>EnumValueDescriptorProto</code></h2>

Describes a value within an enum.

| Fields | |
|---|---|
|`name`||
|`number`||
|`options`||

<h2 id="M_ServiceDescriptorProto">Message <code>ServiceDescriptorProto</code></h2>

Describes a service.

| Fields | |
|---|---|
|`name`||
|`method`||
|`options`||

<h2 id="M_MethodDescriptorProto">Message <code>MethodDescriptorProto</code></h2>

Describes a method of a service.

| Fields | |
|---|---|
|`name`||
|`input_type`|Input and output type names.  These are resolved in the same way as<br/>FieldDescriptorProto.type_name, but must refer to a message type.|
|`output_type`||
|`options`||
|`client_streaming`|Identifies if client streams multiple client messages|
|`server_streaming`|Identifies if server streams multiple server messages|

<h2 id="M_FileOptions">Message <code>FileOptions</code></h2>



| Fields | |
|---|---|
|`java_package`|Sets the Java package where classes generated from this .proto will be<br/>placed.  By default, the proto package is used, but this is often<br/>inappropriate because proto packages do not normally start with backwards<br/>domain names.|
|`java_outer_classname`|Controls the name of the wrapper Java class generated for the .proto file.<br/>That class will always contain the .proto file's getDescriptor() method as<br/>well as any top-level extensions defined in the .proto file.<br/>If java_multiple_files is disabled, then all the other classes from the<br/>.proto file will be nested inside the single wrapper outer class.|
|`java_multiple_files`|If enabled, then the Java code generator will generate a separate .java<br/>file for each top-level message, enum, and service defined in the .proto<br/>file.  Thus, these types will *not* be nested inside the wrapper class<br/>named by java_outer_classname.  However, the wrapper class will still be<br/>generated to contain the file's getDescriptor() method as well as any<br/>top-level extensions defined in the file.|
|`java_generate_equals_and_hash`|This option does nothing.|
|`java_string_check_utf8`|A proto2 file can set this to true to opt in to UTF-8 checking for Java,<br/>which will throw an exception if invalid UTF-8 is parsed from the wire or<br/>assigned to a string field.<br/><br/>TODO: clarify exactly what kinds of field types this option<br/>applies to, and update these docs accordingly.<br/><br/>Proto3 files already perform these checks. Setting the option explicitly to<br/>false has no effect: it cannot be used to opt proto3 files out of UTF-8<br/>checks.|
|`optimize_for`||
|`go_package`|Sets the Go package where structs generated from this .proto will be<br/>placed. If omitted, the Go package will be derived from the following:<br/>- The basename of the package import path, if provided.<br/>- Otherwise, the package statement in the .proto file, if present.<br/>- Otherwise, the basename of the .proto file, without extension.|
|`cc_generic_services`|Should generic services be generated in each language?  "Generic" services<br/>are not specific to any particular RPC system.  They are generated by the<br/>main code generators in each language (without additional plugins).<br/>Generic services were the only kind of service generation supported by<br/>early versions of google.protobuf.<br/><br/>Generic services are now considered deprecated in favor of using plugins<br/>that generate code specific to your particular RPC system.  Therefore,<br/>these default to false.  Old code which depends on generic services should<br/>explicitly set them to true.|
|`java_generic_services`||
|`py_generic_services`||
|`deprecated`|Is this file deprecated?<br/>Depending on the target platform, this can emit Deprecated annotations<br/>for everything in the file, or it will be completely ignored; in the very<br/>least, this is a formalization for deprecating files.|
|`cc_enable_arenas`|Enables the use of arenas for the proto messages in this file. This applies<br/>only to generated classes for C++.|
|`objc_class_prefix`|Sets the objective c class prefix which is prepended to all objective c<br/>generated classes from this .proto. There is no default.|
|`csharp_namespace`|Namespace for generated classes; defaults to the package.|
|`swift_prefix`|By default Swift generators will take the proto package and CamelCase it<br/>replacing '.' with underscore and use that to prefix the types/symbols<br/>defined. When this options is provided, they will use this value instead<br/>to prefix the types/symbols defined.|
|`php_class_prefix`|Sets the php class prefix which is prepended to all php generated classes<br/>from this .proto. Default is empty.|
|`php_namespace`|Use this option to change the namespace of php generated classes. Default<br/>is empty. When this option is empty, the package name will be used for<br/>determining the namespace.|
|`php_metadata_namespace`|Use this option to change the namespace of php generated metadata classes.<br/>Default is empty. When this option is empty, the proto file name will be<br/>used for determining the namespace.|
|`ruby_package`|Use this option to change the package of ruby generated classes. Default<br/>is empty. When this option is not set, the package name will be used for<br/>determining the ruby package.|
|`features`|Any features defined in the specific edition.<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|
|`uninterpreted_option`|The parser stores options it doesn't recognize here.<br/>See the documentation for the "Options" section above.|

<h2 id="M_MessageOptions">Message <code>MessageOptions</code></h2>



| Fields | |
|---|---|
|`message_set_wire_format`|Set true to use the old proto1 MessageSet wire format for extensions.<br/>This is provided for backwards-compatibility with the MessageSet wire<br/>format.  You should not use this for any other reason:  It's less<br/>efficient, has fewer features, and is more complicated.<br/><br/>The message must be defined exactly as follows:<br/>message Foo {<br/>option message_set_wire_format = true;<br/>extensions 4 to max;<br/>}<br/>Note that the message cannot have any defined fields; MessageSets only<br/>have extensions.<br/><br/>All extensions of your type must be singular messages; e.g. they cannot<br/>be int32s, enums, or repeated messages.<br/><br/>Because this is an option, the above two restrictions are not enforced by<br/>the protocol compiler.|
|`no_standard_descriptor_accessor`|Disables the generation of the standard "descriptor()" accessor, which can<br/>conflict with a field of the same name.  This is meant to make migration<br/>from proto1 easier; new code should avoid fields named "descriptor".|
|`deprecated`|Is this message deprecated?<br/>Depending on the target platform, this can emit Deprecated annotations<br/>for the message, or it will be completely ignored; in the very least,<br/>this is a formalization for deprecating messages.|
|`map_entry`|Whether the message is an automatically generated map entry type for the<br/>maps field.<br/><br/>For maps fields:<br/>map<KeyType, ValueType> map_field = 1;<br/>The parsed descriptor looks like:<br/>message MapFieldEntry {<br/>option map_entry = true;<br/>optional KeyType key = 1;<br/>optional ValueType value = 2;<br/>}<br/>repeated MapFieldEntry map_field = 1;<br/><br/>Implementations may choose not to generate the map_entry=true message, but<br/>use a native map in the target language to hold the keys and values.<br/>The reflection APIs in such implementations still need to work as<br/>if the field is a repeated message field.<br/><br/>NOTE: Do not set the option in .proto files. Always use the maps syntax<br/>instead. The option should only be implicitly set by the proto compiler<br/>parser.|
|`deprecated_legacy_json_field_conflicts`|Enable the legacy handling of JSON field name conflicts.  This lowercases<br/>and strips underscored from the fields before comparison in proto3 only.<br/>The new behavior takes `json_name` into account and applies to proto2 as<br/>well.<br/><br/>This should only be used as a temporary measure against broken builds due<br/>to the change in behavior for JSON field name conflicts.<br/><br/>TODO This is legacy behavior we plan to remove once downstream<br/>teams have had time to migrate.|
|`features`|Any features defined in the specific edition.<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|
|`uninterpreted_option`|The parser stores options it doesn't recognize here. See above.|

<h2 id="M_FieldOptions">Message <code>FieldOptions</code></h2>



| Fields | |
|---|---|
|`ctype`|NOTE: ctype is deprecated. Use `features.(pb.cpp).string_type` instead.<br/>The ctype option instructs the C++ code generator to use a different<br/>representation of the field than it normally would.  See the specific<br/>options below.  This option is only implemented to support use of<br/>[ctype=CORD] and [ctype=STRING] (the default) on non-repeated fields of<br/>type "bytes" in the open source release.<br/>TODO: make ctype actually deprecated.|
|`packed`|The packed option can be enabled for repeated primitive fields to enable<br/>a more efficient representation on the wire. Rather than repeatedly<br/>writing the tag and type for each element, the entire array is encoded as<br/>a single length-delimited blob. In proto3, only explicit setting it to<br/>false will avoid using packed encoding.  This option is prohibited in<br/>Editions, but the `repeated_field_encoding` feature can be used to control<br/>the behavior.|
|`jstype`|The jstype option determines the JavaScript type used for values of the<br/>field.  The option is permitted only for 64 bit integral and fixed types<br/>(int64, uint64, sint64, fixed64, sfixed64).  A field with jstype JS_STRING<br/>is represented as JavaScript string, which avoids loss of precision that<br/>can happen when a large value is converted to a floating point JavaScript.<br/>Specifying JS_NUMBER for the jstype causes the generated JavaScript code to<br/>use the JavaScript "number" type.  The behavior of the default option<br/>JS_NORMAL is implementation dependent.<br/><br/>This option is an enum to permit additional types to be added, e.g.<br/>goog.math.Integer.|
|`lazy`|Should this field be parsed lazily?  Lazy applies only to message-type<br/>fields.  It means that when the outer message is initially parsed, the<br/>inner message's contents will not be parsed but instead stored in encoded<br/>form.  The inner message will actually be parsed when it is first accessed.<br/><br/>This is only a hint.  Implementations are free to choose whether to use<br/>eager or lazy parsing regardless of the value of this option.  However,<br/>setting this option true suggests that the protocol author believes that<br/>using lazy parsing on this field is worth the additional bookkeeping<br/>overhead typically needed to implement it.<br/><br/>This option does not affect the public interface of any generated code;<br/>all method signatures remain the same.  Furthermore, thread-safety of the<br/>interface is not affected by this option; const methods remain safe to<br/>call from multiple threads concurrently, while non-const methods continue<br/>to require exclusive access.<br/><br/>Note that lazy message fields are still eagerly verified to check<br/>ill-formed wireformat or missing required fields. Calling IsInitialized()<br/>on the outer message would fail if the inner message has missing required<br/>fields. Failed verification would result in parsing failure (except when<br/>uninitialized messages are acceptable).|
|`unverified_lazy`|unverified_lazy does no correctness checks on the byte stream. This should<br/>only be used where lazy with verification is prohibitive for performance<br/>reasons.|
|`deprecated`|Is this field deprecated?<br/>Depending on the target platform, this can emit Deprecated annotations<br/>for accessors, or it will be completely ignored; in the very least, this<br/>is a formalization for deprecating fields.|
|`weak`|For Google-internal migration only. Do not use.|
|`debug_redact`|Indicate that the field value should not be printed out when using debug<br/>formats, e.g. when the field contains sensitive credentials.|
|`retention`||
|`targets`||
|`edition_defaults`||
|`features`|Any features defined in the specific edition.<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|
|`feature_support`||
|`uninterpreted_option`|The parser stores options it doesn't recognize here. See above.|

<h2 id="M_OneofOptions">Message <code>OneofOptions</code></h2>



| Fields | |
|---|---|
|`features`|Any features defined in the specific edition.<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|
|`uninterpreted_option`|The parser stores options it doesn't recognize here. See above.|

<h2 id="M_EnumOptions">Message <code>EnumOptions</code></h2>



| Fields | |
|---|---|
|`allow_alias`|Set this option to true to allow mapping different tag names to the same<br/>value.|
|`deprecated`|Is this enum deprecated?<br/>Depending on the target platform, this can emit Deprecated annotations<br/>for the enum, or it will be completely ignored; in the very least, this<br/>is a formalization for deprecating enums.|
|`deprecated_legacy_json_field_conflicts`|Enable the legacy handling of JSON field name conflicts.  This lowercases<br/>and strips underscored from the fields before comparison in proto3 only.<br/>The new behavior takes `json_name` into account and applies to proto2 as<br/>well.<br/>TODO Remove this legacy behavior once downstream teams have<br/>had time to migrate.|
|`features`|Any features defined in the specific edition.<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|
|`uninterpreted_option`|The parser stores options it doesn't recognize here. See above.|

<h2 id="M_EnumValueOptions">Message <code>EnumValueOptions</code></h2>



| Fields | |
|---|---|
|`deprecated`|Is this enum value deprecated?<br/>Depending on the target platform, this can emit Deprecated annotations<br/>for the enum value, or it will be completely ignored; in the very least,<br/>this is a formalization for deprecating enum values.|
|`features`|Any features defined in the specific edition.<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|
|`debug_redact`|Indicate that fields annotated with this enum value should not be printed<br/>out when using debug formats, e.g. when the field contains sensitive<br/>credentials.|
|`feature_support`|Information about the support window of a feature value.|
|`uninterpreted_option`|The parser stores options it doesn't recognize here. See above.|

<h2 id="M_ServiceOptions">Message <code>ServiceOptions</code></h2>



| Fields | |
|---|---|
|`features`|Any features defined in the specific edition.<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|
|`deprecated`|Is this service deprecated?<br/>Depending on the target platform, this can emit Deprecated annotations<br/>for the service, or it will be completely ignored; in the very least,<br/>this is a formalization for deprecating services.|
|`uninterpreted_option`|The parser stores options it doesn't recognize here. See above.|

<h2 id="M_MethodOptions">Message <code>MethodOptions</code></h2>



| Fields | |
|---|---|
|`deprecated`|Is this method deprecated?<br/>Depending on the target platform, this can emit Deprecated annotations<br/>for the method, or it will be completely ignored; in the very least,<br/>this is a formalization for deprecating methods.|
|`idempotency_level`||
|`features`|Any features defined in the specific edition.<br/>WARNING: This field should only be used by protobuf plugins or special<br/>cases like the proto compiler. Other uses are discouraged and<br/>developers should rely on the protoreflect APIs for their client language.|
|`uninterpreted_option`|The parser stores options it doesn't recognize here. See above.|

<h2 id="M_UninterpretedOption">Message <code>UninterpretedOption</code></h2>

A message representing a option the parser does not recognize. This only
 appears in options protos created by the compiler::Parser class.
 DescriptorPool resolves these when building Descriptor objects. Therefore,
 options protos in descriptor objects (e.g. returned by Descriptor::options(),
 or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
 in them.

| Fields | |
|---|---|
|`name`||
|`identifier_value`|The value of the uninterpreted option, in whatever type the tokenizer<br/>identified it as during parsing. Exactly one of these should be set.|
|`positive_int_value`||
|`negative_int_value`||
|`double_value`||
|`string_value`||
|`aggregate_value`||

<h2 id="M_FeatureSet">Message <code>FeatureSet</code></h2>

TODO Enums in C++ gencode (and potentially other languages) are
 not well scoped.  This means that each of the feature enums below can clash
 with each other.  The short names we've chosen maximize call-site
 readability, but leave us very open to this scenario.  A future feature will
 be designed and implemented to handle this, hopefully before we ever hit a
 conflict here.

| Fields | |
|---|---|
|`field_presence`||
|`enum_type`||
|`repeated_field_encoding`||
|`utf8_validation`||
|`message_encoding`||
|`json_format`||
|`enforce_naming_style`||
|`default_symbol_visibility`||

<h2 id="M_FeatureSetDefaults">Message <code>FeatureSetDefaults</code></h2>

A compiled specification for the defaults of a set of features.  These
 messages are generated from FeatureSet extensions and can be used to seed
 feature resolution. The resolution with this object becomes a simple search
 for the closest matching edition, followed by proto merges.

| Fields | |
|---|---|
|`defaults`||
|`minimum_edition`|The minimum supported edition (inclusive) when this was constructed.<br/>Editions before this will not have defaults.|
|`maximum_edition`|The maximum known edition (inclusive) when this was constructed. Editions<br/>after this will not have reliable defaults.|

<h2 id="M_SourceCodeInfo">Message <code>SourceCodeInfo</code></h2>

Encapsulates information about the original source file from which a
 FileDescriptorProto was generated.

| Fields | |
|---|---|
|`location`|A Location identifies a piece of source code in a .proto file which<br/>corresponds to a particular definition.  This information is intended<br/>to be useful to IDEs, code indexers, documentation generators, and similar<br/>tools.<br/><br/>For example, say we have a file like:<br/>message Foo {<br/>optional string foo = 1;<br/>}<br/>Let's look at just the field definition:<br/>optional string foo = 1;<br/>^       ^^     ^^  ^  ^^^<br/>a       bc     de  f  ghi<br/>We have the following locations:<br/>span   path               represents<br/>[a,i)  [ 4, 0, 2, 0 ]     The whole field definition.<br/>[a,b)  [ 4, 0, 2, 0, 4 ]  The label (optional).<br/>[c,d)  [ 4, 0, 2, 0, 5 ]  The type (string).<br/>[e,f)  [ 4, 0, 2, 0, 1 ]  The name (foo).<br/>[g,h)  [ 4, 0, 2, 0, 3 ]  The number (1).<br/><br/>Notes:<br/>- A location may refer to a repeated field itself (i.e. not to any<br/>particular index within it).  This is used whenever a set of elements are<br/>logically enclosed in a single code segment.  For example, an entire<br/>extend block (possibly containing multiple extension definitions) will<br/>have an outer location whose path refers to the "extensions" repeated<br/>field without an index.<br/>- Multiple locations may have the same path.  This happens when a single<br/>logical declaration is spread out across multiple places.  The most<br/>obvious example is the "extend" block again -- there may be multiple<br/>extend blocks in the same scope, each of which will have the same path.<br/>- A location's span is not always a subset of its parent's span.  For<br/>example, the "extendee" of an extension declaration appears at the<br/>beginning of the "extend" block and is shared by all extensions within<br/>the block.<br/>- Just because a location's span is a subset of some other location's span<br/>does not mean that it is a descendant.  For example, a "group" defines<br/>both a type and a field in a single declaration.  Thus, the locations<br/>corresponding to the type and field and their components will overlap.<br/>- Code which tries to interpret locations should probably be designed to<br/>ignore those that it doesn't understand, as more types of locations could<br/>be recorded in the future.|

<h2 id="M_GeneratedCodeInfo">Message <code>GeneratedCodeInfo</code></h2>

Describes the relationship between generated code and its original source
 file. A GeneratedCodeInfo message is associated with only one generated
 source file, but may contain references to different source .proto files.

| Fields | |
|---|---|
|`annotation`|An Annotation connects some span of text in generated code to an element<br/>of its generating .proto file.|

<h2 id="M_Timestamp">Message <code>Timestamp</code></h2>

A Timestamp represents a point in time independent of any time zone or local
 calendar, encoded as a count of seconds and fractions of seconds at
 nanosecond resolution. The count is relative to an epoch at UTC midnight on
 January 1, 1970, in the proleptic Gregorian calendar which extends the
 Gregorian calendar backwards to year one.

 All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
 second table is needed for interpretation, using a [24-hour linear
 smear](https://developers.google.com/time/smear).

 The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
 restricting to that range, we ensure that we can convert to and from [RFC
 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.

 # Examples

 Example 1: Compute Timestamp from POSIX `time()`.

     Timestamp timestamp;
     timestamp.set_seconds(time(NULL));
     timestamp.set_nanos(0);

 Example 2: Compute Timestamp from POSIX `gettimeofday()`.

     struct timeval tv;
     gettimeofday(&tv, NULL);

     Timestamp timestamp;
     timestamp.set_seconds(tv.tv_sec);
     timestamp.set_nanos(tv.tv_usec * 1000);

 Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

     FILETIME ft;
     GetSystemTimeAsFileTime(&ft);
     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;

     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
     Timestamp timestamp;
     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

 Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

     long millis = System.currentTimeMillis();

     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
         .setNanos((int) ((millis % 1000) * 1000000)).build();

 Example 5: Compute Timestamp from Java `Instant.now()`.

     Instant now = Instant.now();

     Timestamp timestamp =
         Timestamp.newBuilder().setSeconds(now.getEpochSecond())
             .setNanos(now.getNano()).build();

 Example 6: Compute Timestamp from current time in Python.

     timestamp = Timestamp()
     timestamp.GetCurrentTime()

 # JSON Mapping

 In JSON format, the Timestamp type is encoded as a string in the
 [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
 format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
 where {year} is always expressed using four digits while {month}, {day},
 {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
 seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
 are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
 is required. A proto3 JSON serializer should always use UTC (as indicated by
 "Z") when printing the Timestamp type and a proto3 JSON parser should be
 able to accept both UTC and other timezones (as indicated by an offset).

 For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
 01:30 UTC on January 15, 2017.

 In JavaScript, one can convert a Date object to this format using the
 standard
 [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
 method. In Python, a standard `datetime.datetime` object can be converted
 to this format using
 [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
 the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
 the Joda Time's [`ISODateTimeFormat.dateTime()`](
 http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()
 ) to obtain a formatter capable of generating timestamps in this format.

| Fields | |
|---|---|
|`seconds`|Represents seconds of UTC time since Unix epoch<br/>1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to<br/>9999-12-31T23:59:59Z inclusive.|
|`nanos`|Non-negative fractions of a second at nanosecond resolution. Negative<br/>second values with fractions must still have non-negative nanos values<br/>that count forward in time. Must be from 0 to 999,999,999<br/>inclusive.|

<h2 id="M_Duration">Message <code>Duration</code></h2>

A Duration represents a signed, fixed-length span of time represented
 as a count of seconds and fractions of seconds at nanosecond
 resolution. It is independent of any calendar and concepts like "day"
 or "month". It is related to Timestamp in that the difference between
 two Timestamp values is a Duration and it can be added or subtracted
 from a Timestamp. Range is approximately +-10,000 years.

 # Examples

 Example 1: Compute Duration from two Timestamps in pseudo code.

     Timestamp start = ...;
     Timestamp end = ...;
     Duration duration = ...;

     duration.seconds = end.seconds - start.seconds;
     duration.nanos = end.nanos - start.nanos;

     if (duration.seconds < 0 && duration.nanos > 0) {
       duration.seconds += 1;
       duration.nanos -= 1000000000;
     } else if (duration.seconds > 0 && duration.nanos < 0) {
       duration.seconds -= 1;
       duration.nanos += 1000000000;
     }

 Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.

     Timestamp start = ...;
     Duration duration = ...;
     Timestamp end = ...;

     end.seconds = start.seconds + duration.seconds;
     end.nanos = start.nanos + duration.nanos;

     if (end.nanos < 0) {
       end.seconds -= 1;
       end.nanos += 1000000000;
     } else if (end.nanos >= 1000000000) {
       end.seconds += 1;
       end.nanos -= 1000000000;
     }

 Example 3: Compute Duration from datetime.timedelta in Python.

     td = datetime.timedelta(days=3, minutes=10)
     duration = Duration()
     duration.FromTimedelta(td)

 # JSON Mapping

 In JSON format, the Duration type is encoded as a string rather than an
 object, where the string ends in the suffix "s" (indicating seconds) and
 is preceded by the number of seconds, with nanoseconds expressed as
 fractional seconds. For example, 3 seconds with 0 nanoseconds should be
 encoded in JSON format as "3s", while 3 seconds and 1 nanosecond should
 be expressed in JSON format as "3.000000001s", and 3 seconds and 1
 microsecond should be expressed in JSON format as "3.000001s".

| Fields | |
|---|---|
|`seconds`|Signed seconds of the span of time. Must be from -315,576,000,000<br/>to +315,576,000,000 inclusive. Note: these bounds are computed from:<br/>60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years|
|`nanos`|Signed fractions of a second at nanosecond resolution of the span<br/>of time. Durations less than one second are represented with a 0<br/>`seconds` field and a positive or negative `nanos` field. For durations<br/>of one second or more, a non-zero value for the `nanos` field must be<br/>of the same sign as the `seconds` field. Must be from -999,999,999<br/>to +999,999,999 inclusive.|

<h2 id="M_Any">Message <code>Any</code></h2>

`Any` contains an arbitrary serialized protocol buffer message along with a
 URL that describes the type of the serialized message.

 Protobuf library provides support to pack/unpack Any values in the form
 of utility functions or additional generated methods of the Any type.

 Example 1: Pack and unpack a message in C++.

     Foo foo = ...;
     Any any;
     any.PackFrom(foo);
     ...
     if (any.UnpackTo(&foo)) {
       ...
     }

 Example 2: Pack and unpack a message in Java.

     Foo foo = ...;
     Any any = Any.pack(foo);
     ...
     if (any.is(Foo.class)) {
       foo = any.unpack(Foo.class);
     }
     // or ...
     if (any.isSameTypeAs(Foo.getDefaultInstance())) {
       foo = any.unpack(Foo.getDefaultInstance());
     }

  Example 3: Pack and unpack a message in Python.

     foo = Foo(...)
     any = Any()
     any.Pack(foo)
     ...
     if any.Is(Foo.DESCRIPTOR):
       any.Unpack(foo)
       ...

  Example 4: Pack and unpack a message in Go

      foo := &pb.Foo{...}
      any, err := anypb.New(foo)
      if err != nil {
        ...
      }
      ...
      foo := &pb.Foo{}
      if err := any.UnmarshalTo(foo); err != nil {
        ...
      }

 The pack methods provided by protobuf library will by default use
 'type.googleapis.com/full.type.name' as the type URL and the unpack
 methods only use the fully qualified type name after the last '/'
 in the type URL, for example "foo.bar.com/x/y.z" will yield type
 name "y.z".

 JSON
 ====
 The JSON representation of an `Any` value uses the regular
 representation of the deserialized, embedded message, with an
 additional field `@type` which contains the type URL. Example:

     package google.profile;
     message Person {
       string first_name = 1;
       string last_name = 2;
     }

     {
       "@type": "type.googleapis.com/google.profile.Person",
       "firstName": <string>,
       "lastName": <string>
     }

 If the embedded message type is well-known and has a custom JSON
 representation, that representation will be embedded adding a field
 `value` which holds the custom JSON in addition to the `@type`
 field. Example (for message [google.protobuf.Duration][]):

     {
       "@type": "type.googleapis.com/google.protobuf.Duration",
       "value": "1.212s"
     }

| Fields | |
|---|---|
|`type_url`|A URL/resource name that uniquely identifies the type of the serialized<br/>protocol buffer message. This string must contain at least<br/>one "/" character. The last segment of the URL's path must represent<br/>the fully qualified name of the type (as in<br/>`path/google.protobuf.Duration`). The name should be in a canonical form<br/>(e.g., leading "." is not accepted).<br/><br/>In practice, teams usually precompile into the binary all types that they<br/>expect it to use in the context of Any. However, for URLs which use the<br/>scheme `http`, `https`, or no scheme, one can optionally set up a type<br/>server that maps type URLs to message definitions as follows:<br/><br/>* If no scheme is provided, `https` is assumed.<br/>* An HTTP GET on the URL must yield a [google.protobuf.Type][]<br/>value in binary format, or produce an error.<br/>* Applications are allowed to cache lookup results based on the<br/>URL, or have them precompiled into a binary to avoid any<br/>lookup. Therefore, binary compatibility needs to be preserved<br/>on changes to types. (Use versioned type names to manage<br/>breaking changes.)<br/><br/>Note: this functionality is not currently available in the official<br/>protobuf release, and it is not used for type URLs beginning with<br/>type.googleapis.com. As of May 2023, there are no widely used type server<br/>implementations and no plans to implement one.<br/><br/>Schemes other than `http`, `https` (or the empty scheme) might be<br/>used with implementation specific semantics.|
|`value`|Must be a valid serialized protocol buffer of the above specified type.|

<h2 id="M_Empty">Message <code>Empty</code></h2>

A generic empty message that you can re-use to avoid defining duplicated
 empty messages in your APIs. A typical example is to use it as the request
 or the response type of an API method. For instance:

     service Foo {
       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
     }

| Fields | |
|---|---|

<h2 id="M_FieldMask">Message <code>FieldMask</code></h2>

`FieldMask` represents a set of symbolic field paths, for example:

     paths: "f.a"
     paths: "f.b.d"

 Here `f` represents a field in some root message, `a` and `b`
 fields in the message found in `f`, and `d` a field found in the
 message in `f.b`.

 Field masks are used to specify a subset of fields that should be
 returned by a get operation or modified by an update operation.
 Field masks also have a custom JSON encoding (see below).

 # Field Masks in Projections

 When used in the context of a projection, a response message or
 sub-message is filtered by the API to only contain those fields as
 specified in the mask. For example, if the mask in the previous
 example is applied to a response message as follows:

     f {
       a : 22
       b {
         d : 1
         x : 2
       }
       y : 13
     }
     z: 8

 The result will not contain specific values for fields x,y and z
 (their value will be set to the default, and omitted in proto text
 output):


     f {
       a : 22
       b {
         d : 1
       }
     }

 A repeated field is not allowed except at the last position of a
 paths string.

 If a FieldMask object is not present in a get operation, the
 operation applies to all fields (as if a FieldMask of all fields
 had been specified).

 Note that a field mask does not necessarily apply to the
 top-level response message. In case of a REST get operation, the
 field mask applies directly to the response, but in case of a REST
 list operation, the mask instead applies to each individual message
 in the returned resource list. In case of a REST custom method,
 other definitions may be used. Where the mask applies will be
 clearly documented together with its declaration in the API.  In
 any case, the effect on the returned resource/resources is required
 behavior for APIs.

 # Field Masks in Update Operations

 A field mask in update operations specifies which fields of the
 targeted resource are going to be updated. The API is required
 to only change the values of the fields as specified in the mask
 and leave the others untouched. If a resource is passed in to
 describe the updated values, the API ignores the values of all
 fields not covered by the mask.

 If a repeated field is specified for an update operation, new values will
 be appended to the existing repeated field in the target resource. Note that
 a repeated field is only allowed in the last position of a `paths` string.

 If a sub-message is specified in the last position of the field mask for an
 update operation, then new value will be merged into the existing sub-message
 in the target resource.

 For example, given the target message:

     f {
       b {
         d: 1
         x: 2
       }
       c: [1]
     }

 And an update message:

     f {
       b {
         d: 10
       }
       c: [2]
     }

 then if the field mask is:

  paths: ["f.b", "f.c"]

 then the result will be:

     f {
       b {
         d: 10
         x: 2
       }
       c: [1, 2]
     }

 An implementation may provide options to override this default behavior for
 repeated and message fields.

 In order to reset a field's value to the default, the field must
 be in the mask and set to the default value in the provided resource.
 Hence, in order to reset all fields of a resource, provide a default
 instance of the resource and set all fields in the mask, or do
 not provide a mask as described below.

 If a field mask is not present on update, the operation applies to
 all fields (as if a field mask of all fields has been specified).
 Note that in the presence of schema evolution, this may mean that
 fields the client does not know and has therefore not filled into
 the request will be reset to their default. If this is unwanted
 behavior, a specific service may require a client to always specify
 a field mask, producing an error if not.

 As with get operations, the location of the resource which
 describes the updated values in the request message depends on the
 operation kind. In any case, the effect of the field mask is
 required to be honored by the API.

 ## Considerations for HTTP REST

 The HTTP kind of an update operation which uses a field mask must
 be set to PATCH instead of PUT in order to satisfy HTTP semantics
 (PUT must only be used for full updates).

 # JSON Encoding of Field Masks

 In JSON, a field mask is encoded as a single string where paths are
 separated by a comma. Fields name in each path are converted
 to/from lower-camel naming conventions.

 As an example, consider the following message declarations:

     message Profile {
       User user = 1;
       Photo photo = 2;
     }
     message User {
       string display_name = 1;
       string address = 2;
     }

 In proto a field mask for `Profile` may look as such:

     mask {
       paths: "user.display_name"
       paths: "photo"
     }

 In JSON, the same mask is represented as below:

     {
       mask: "user.displayName,photo"
     }

 # Field Masks and Oneof Fields

 Field masks treat fields in oneofs just as regular fields. Consider the
 following message:

     message SampleMessage {
       oneof test_oneof {
         string name = 4;
         SubMessage sub_message = 9;
       }
     }

 The field mask can be:

     mask {
       paths: "name"
     }

 Or:

     mask {
       paths: "sub_message"
     }

 Note that oneof type names ("test_oneof" in this case) cannot be used in
 paths.

 ## Field Mask Verification

 The implementation of any API method which has a FieldMask type field in the
 request should verify the included field paths, and return an
 `INVALID_ARGUMENT` error if any path is unmappable.

| Fields | |
|---|---|
|`paths`|The set of field mask paths.|

