package walk

import "google.golang.org/protobuf/types/descriptorpb"

// file-level field numbers (from descriptor.proto)
const (
	fileMessageType = 4
	fileEnumType    = 5
	fileService     = 6
)

// message-level field numbers
const (
	msgField      = 2
	msgNestedType = 3
	msgEnumType   = 4
)

// service-level field numbers
const (
	svcMethod = 2
)

type Visitor struct {
	OnMessage func(path []int32, m *descriptorpb.DescriptorProto)
	OnField   func(path []int32, f *descriptorpb.FieldDescriptorProto)
	OnEnum    func(path []int32, e *descriptorpb.EnumDescriptorProto)
	OnService func(path []int32, s *descriptorpb.ServiceDescriptorProto)
	OnMethod  func(path []int32, m *descriptorpb.MethodDescriptorProto)
}

func WalkFile(f *descriptorpb.FileDescriptorProto, v Visitor) {
	// top-level messages
	for i, m := range f.GetMessageType() {
		p := []int32{fileMessageType, int32(i)}
		walkMessage(m, p, v)
	}

	// top-level enums
	for i, e := range f.GetEnumType() {
		p := []int32{fileEnumType, int32(i)}
		if v.OnEnum != nil {
			v.OnEnum(p, e)
		}
	}

	// services + methods
	for i, s := range f.GetService() {
		sp := []int32{fileService, int32(i)}
		if v.OnService != nil {
			v.OnService(sp, s)
		}
		for j, m := range s.GetMethod() {
			mp := append(append([]int32{}, sp...), svcMethod, int32(j))
			if v.OnMethod != nil {
				v.OnMethod(mp, m)
			}
		}
	}
}

func walkMessage(m *descriptorpb.DescriptorProto, path []int32, v Visitor) {
	if v.OnMessage != nil {
		v.OnMessage(path, m)
	}

	// fields
	for i, f := range m.GetField() {
		fp := append(append([]int32{}, path...), msgField, int32(i))
		if v.OnField != nil {
			v.OnField(fp, f)
		}
	}

	// nested enums
	for i, e := range m.GetEnumType() {
		ep := append(append([]int32{}, path...), msgEnumType, int32(i))
		if v.OnEnum != nil {
			v.OnEnum(ep, e)
		}
	}

	// nested messages (recursive)
	for i, n := range m.GetNestedType() {
		np := append(append([]int32{}, path...), msgNestedType, int32(i))
		walkMessage(n, np, v)
	}
}
