package models

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

type Package struct {
	Name     string
	Messages []Message
	Services []Service
	Enums    []Enum
}

func NewPackage(name string) Package {
	return Package{
		Name: name,
	}
}

func (p *Package) AddProtoMessage(proto *descriptorpb.DescriptorProto, comments CommentMap, path []int32) {
	p.Messages = append(p.Messages, NewMessage(proto, p, comments, path))
}

func (p *Package) AddProtoService(proto *descriptorpb.ServiceDescriptorProto, comments CommentMap, path []int32) {
	p.Services = append(p.Services, NewService(proto, p, comments, path))
}

func (p *Package) AddProtoEnum(proto *descriptorpb.EnumDescriptorProto, comments CommentMap, path []int32) {
	p.Enums = append(p.Enums, NewEnum(proto, p, nil, comments, path))
}
