package mdgen

import (
	"fmt"
	"strings"

	"github.com/roostmoe/protomd/internal/registry"
	"github.com/roostmoe/protomd/internal/wellknown"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Field struct {
	Name         string
	DocComment   string
	Behaviours   []string
	FieldType    string
	Optional     bool
	Repeated     bool
	ShouldLink   bool
	LinkedType   string // the message name to link to (empty if not a message)
	WellKnownURL string // external URL if this is a well-known type
	IsWellKnown  bool   // true if this is a well-known type
	OneofIndex   *int32 // index of the oneof this field belongs to (nil if not part of a oneof)
}

func NewField(
	reg *registry.Registry,
	pb *descriptorpb.FieldDescriptorProto,
	fileName string,
	path []int32,
) Field {
	field := Field{
		Name:       pb.GetName(),
		Behaviours: []string{},
		OneofIndex: pb.OneofIndex, // track oneof membership
	}

	// get comment from registry
	if comments, ok := reg.Comments[fileName]; ok {
		if loc, ok := comments[pathKey(path)]; ok {
			field.DocComment = cleanComment(loc.GetLeadingComments())
		}
	}

	field.FieldType = getFieldType(pb)

	// check if this field references another message in the same package
	if kind := pb.GetType(); kind == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
		typeName := pb.GetTypeName()

		// check if it's a well-known type first
		if wellknown.IsWellKnown(typeName) {
			field.IsWellKnown = true
			field.WellKnownURL = wellknown.GetDocURL(typeName)
			// still set LinkedType for the display name
			parts := strings.Split(typeName, ".")
			if len(parts) > 0 {
				field.LinkedType = parts[len(parts)-1]
			}
		} else {
			// type names are like ".hello.MyResource" - extract just the message name
			parts := strings.Split(typeName, ".")
			if len(parts) > 0 {
				field.LinkedType = parts[len(parts)-1]
			}
		}
	}

	if isRepeated := pb.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED; isRepeated {
		field.FieldType = fmt.Sprintf("[]%s", field.FieldType)
	}

	// check for field behavior annotations
	if options := pb.GetOptions(); options != nil {
		if proto.HasExtension(options, annotations.E_FieldBehavior) {
			behaviors := proto.GetExtension(options, annotations.E_FieldBehavior).([]annotations.FieldBehavior)
			for _, fb := range behaviors {
				field.Behaviours = append(field.Behaviours, fb.String())
			}
		}
	}

	return field
}

func getFieldType(pb *descriptorpb.FieldDescriptorProto) string {
	kind := pb.GetType()

	// for message and enum types, return the type name
	if kind == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE ||
		kind == descriptorpb.FieldDescriptorProto_TYPE_ENUM {
		typeName := pb.GetTypeName()
		// strip leading dot from type names like ".google.protobuf.Timestamp"
		return strings.TrimPrefix(typeName, ".")
	}

	// map scalar types to their proto names
	scalarTypes := map[descriptorpb.FieldDescriptorProto_Type]string{
		descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:   "double",
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT:    "float",
		descriptorpb.FieldDescriptorProto_TYPE_INT64:    "int64",
		descriptorpb.FieldDescriptorProto_TYPE_UINT64:   "uint64",
		descriptorpb.FieldDescriptorProto_TYPE_INT32:    "int32",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64:  "fixed64",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32:  "fixed32",
		descriptorpb.FieldDescriptorProto_TYPE_BOOL:     "bool",
		descriptorpb.FieldDescriptorProto_TYPE_STRING:   "string",
		descriptorpb.FieldDescriptorProto_TYPE_BYTES:    "bytes",
		descriptorpb.FieldDescriptorProto_TYPE_UINT32:   "uint32",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: "sfixed32",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: "sfixed64",
		descriptorpb.FieldDescriptorProto_TYPE_SINT32:   "sint32",
		descriptorpb.FieldDescriptorProto_TYPE_SINT64:   "sint64",
	}

	if typeName, ok := scalarTypes[kind]; ok {
		return typeName
	}

	// fallback to enum string if not found
	return kind.String()
}
