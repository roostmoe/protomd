package models

import "google.golang.org/protobuf/types/descriptorpb"

type Service struct {
	Name        string
	Description string
	Proto       *descriptorpb.ServiceDescriptorProto
	Methods     []Method
}

func NewService(proto *descriptorpb.ServiceDescriptorProto, pkg *Package, comments CommentMap, path []int32) Service {
	service := Service{
		Name:        proto.GetName(),
		Description: getComment(comments, path),
		Proto:       proto,
	}

	for idx, method := range proto.GetMethod() {
		methodPath := append(path, 2, int32(idx))
		service.Methods = append(service.Methods, NewMethod(method, &service, comments, methodPath))
	}

	return service
}
