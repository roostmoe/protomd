package models

// WellKnownTypes maps fully-qualified protobuf type names to their documentation URLs
var WellKnownTypes = map[string]string{
	// google.protobuf well-known types
	".google.protobuf.Timestamp":   "https://protobuf.dev/reference/protobuf/google.protobuf/#timestamp",
	".google.protobuf.Duration":    "https://protobuf.dev/reference/protobuf/google.protobuf/#duration",
	".google.protobuf.Empty":       "https://protobuf.dev/reference/protobuf/google.protobuf/#empty",
	".google.protobuf.Any":         "https://protobuf.dev/reference/protobuf/google.protobuf/#any",
	".google.protobuf.Struct":      "https://protobuf.dev/reference/protobuf/google.protobuf/#struct",
	".google.protobuf.Value":       "https://protobuf.dev/reference/protobuf/google.protobuf/#value",
	".google.protobuf.ListValue":   "https://protobuf.dev/reference/protobuf/google.protobuf/#listvalue",
	".google.protobuf.NullValue":   "https://protobuf.dev/reference/protobuf/google.protobuf/#nullvalue",
	".google.protobuf.FieldMask":   "https://protobuf.dev/reference/protobuf/google.protobuf/#field-mask",
	".google.protobuf.BoolValue":   "https://protobuf.dev/reference/protobuf/google.protobuf/#boolvalue",
	".google.protobuf.Int32Value":  "https://protobuf.dev/reference/protobuf/google.protobuf/#int32value",
	".google.protobuf.Int64Value":  "https://protobuf.dev/reference/protobuf/google.protobuf/#int64value",
	".google.protobuf.UInt32Value": "https://protobuf.dev/reference/protobuf/google.protobuf/#uint32value",
	".google.protobuf.UInt64Value": "https://protobuf.dev/reference/protobuf/google.protobuf/#uint64value",
	".google.protobuf.FloatValue":  "https://protobuf.dev/reference/protobuf/google.protobuf/#floatvalue",
	".google.protobuf.DoubleValue": "https://protobuf.dev/reference/protobuf/google.protobuf/#doublevalue",
	".google.protobuf.StringValue": "https://protobuf.dev/reference/protobuf/google.protobuf/#stringvalue",
	".google.protobuf.BytesValue":  "https://protobuf.dev/reference/protobuf/google.protobuf/#bytesvalue",
}

// GetDocURL returns the documentation URL for a well-known type.
// Returns empty string if the type is not a well-known type.
func GetDocURL(typeName string) string {
	return WellKnownTypes[typeName]
}

// IsWellKnown checks if a type is a well-known type
func IsWellKnown(typeName string) bool {
	_, ok := WellKnownTypes[typeName]
	return ok
}
