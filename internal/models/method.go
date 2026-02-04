package models

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/types/descriptorpb"
)

type HttpRule struct {
	Method string
	Path   string
}

type Method struct {
	Service     *Service
	Name        string
	Description string
	Proto       *descriptorpb.MethodDescriptorProto

	Input  *Message
	Output *Message
}

func NewMethod(proto *descriptorpb.MethodDescriptorProto, svc *Service, comments CommentMap, path []int32) Method {
	m := Method{
		Service:     svc,
		Name:        proto.GetName(),
		Description: getComment(comments, path),
		Proto:       proto,
	}

	return m
}

func (m *Method) SetInputMessage(msg *Message) {
	m.Input = msg
}

func (m *Method) SetOutputMessage(msg *Message) {
	m.Output = msg
}

func (m *Method) GetHttpRule() *HttpRule {
	opts := m.Proto.GetOptions()
	if opts == nil {
		return nil
	}

	rule, ok := getMethodExtension[annotations.HttpRule](opts, annotations.E_Http)
	if !ok {
		return nil
	}

	httpRule := &HttpRule{}

	switch r := rule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		httpRule.Method = "GET"
		httpRule.Path = r.Get
	case *annotations.HttpRule_Post:
		httpRule.Method = "POST"
		httpRule.Path = r.Post
	case *annotations.HttpRule_Patch:
		httpRule.Method = "PATCH"
		httpRule.Path = r.Patch
	case *annotations.HttpRule_Put:
		httpRule.Method = "PUT"
		httpRule.Path = r.Put
	case *annotations.HttpRule_Delete:
		httpRule.Method = "DELETE"
		httpRule.Path = r.Delete
	case *annotations.HttpRule_Custom:
		httpRule.Method = r.Custom.Kind
		httpRule.Path = r.Custom.Path
	default:
		return nil
	}

	return httpRule
}
