package mdgen

import (
	"io"

	"github.com/roostmoe/protomd/internal/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Service struct {
	Name       string
	DocComment string

	Methods []Method
	HasHttp bool
}

func NewService(
	reg *registry.Registry,
	proto *descriptorpb.ServiceDescriptorProto,
	fileName string,
	path []int32,
) Service {
	s := Service{
		Name:    proto.GetName(),
		Methods: []Method{},
	}

	// get comment from registry
	if comments, ok := reg.Comments[fileName]; ok {
		if loc, ok := comments[pathKey(path)]; ok {
			s.DocComment = cleanComment(loc.GetLeadingComments())
		}
	}

	// build methods
	for i, m := range proto.GetMethod() {
		methodPath := append(append([]int32{}, path...), 2, int32(i)) // 2 = method tag in ServiceDescriptorProto
		method := NewMethod(reg, m, fileName, methodPath)
		s.Methods = append(s.Methods, method)
		if method.HasHttp {
			s.HasHttp = true
		}
	}

	return s
}

func (s Service) WriteGrpc(w io.Writer) error {
	type data struct {
		Service Service
	}

	return tmpl.ExecuteTemplate(w, "grpc-service.md", data{Service: s})
}

func (s Service) WriteRest(w io.Writer) error {
	type data struct {
		Service Service
	}

	return tmpl.ExecuteTemplate(w, "rest-service.md", data{Service: s})
}
