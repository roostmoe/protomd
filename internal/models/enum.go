package models

import "google.golang.org/protobuf/types/descriptorpb"

type Enum struct {
	Package     *Package
	Message     *Message
	Name        string
	Description string
	Proto       *descriptorpb.EnumDescriptorProto
	Values      []EnumValue
}

type EnumValue struct {
	Enum        *Enum
	Name        string
	Value       int32
	Description string
	Proto       *descriptorpb.EnumValueDescriptorProto
}

func NewEnum(proto *descriptorpb.EnumDescriptorProto, pkg *Package, msg *Message, comments CommentMap, path []int32) Enum {
	enum := Enum{
		Package:     pkg,
		Message:     msg,
		Name:        proto.GetName(),
		Description: getComment(comments, path),
		Proto:       proto,
	}

	for idx, value := range proto.GetValue() {
		valuePath := append(path, 2, int32(idx))
		enum.Values = append(enum.Values, NewEnumValue(value, &enum, comments, valuePath))
	}

	return enum
}

func NewEnumValue(proto *descriptorpb.EnumValueDescriptorProto, enum *Enum, comments CommentMap, path []int32) EnumValue {
	return EnumValue{
		Enum:        enum,
		Name:        proto.GetName(),
		Value:       proto.GetNumber(),
		Description: getComment(comments, path),
		Proto:       proto,
	}
}
