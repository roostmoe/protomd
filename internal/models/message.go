package models

import "google.golang.org/protobuf/types/descriptorpb"

type Message struct {
	Package     *Package
	Name        string
	Description string
	Proto       *descriptorpb.DescriptorProto

	Fields   []Field
	Enums    []Enum
	Messages []Message
}

type Field struct {
	Message     *Message
	Name        string
	Description string
	Proto       *descriptorpb.FieldDescriptorProto
}

func NewMessage(proto *descriptorpb.DescriptorProto, pkg *Package, comments CommentMap, path []int32) Message {
	message := Message{
		Package:     pkg,
		Name:        proto.GetName(),
		Description: getComment(comments, path),
		Proto:       proto,
	}

	for idx, field := range proto.GetField() {
		fieldPath := append(path, 2, int32(idx))
		message.Fields = append(message.Fields, NewField(field, &message, comments, fieldPath))
	}

	for idx, enum := range proto.GetEnumType() {
		enumPath := append(path, 2, int32(idx))
		message.Enums = append(message.Enums, NewEnum(enum, nil, &message, comments, enumPath))
	}

	for idx, msg := range proto.GetNestedType() {
		msgPath := append(path, 2, int32(idx))
		message.Messages = append(message.Messages, NewMessage(msg, pkg, comments, msgPath))
	}

	return message
}

func NewField(proto *descriptorpb.FieldDescriptorProto, msg *Message, comments CommentMap, path []int32) Field {
	return Field{
		Message:     msg,
		Name:        proto.GetName(),
		Description: getComment(comments, path),
		Proto:       proto,
	}
}
