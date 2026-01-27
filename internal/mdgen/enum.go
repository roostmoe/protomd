package mdgen

import (
	"github.com/roostmoe/protomd/internal/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type EnumValue struct {
	Name       string
	DocComment string
}

type Enum struct {
	Name       string
	DocComment string
	Options    []EnumValue
}

func NewEnum(
	reg *registry.Registry,
	proto *descriptorpb.EnumDescriptorProto,
	fileName string,
	path []int32,
) Enum {
	m := Enum{
		Name: proto.GetName(),
	}

	// get comment from registry
	if comments, ok := reg.Comments[fileName]; ok {
		if loc, ok := comments[pathKey(path)]; ok {
			m.DocComment = cleanComment(loc.GetLeadingComments())
		}
	}

	// collect enum values
	for i, val := range proto.GetValue() {
		// path for enum value: append field 2 (value field) and index
		valPath := append(append([]int32{}, path...), 2, int32(i))

		enumVal := EnumValue{
			Name: val.GetName(),
		}

		// get value comment from registry
		if comments, ok := reg.Comments[fileName]; ok {
			if loc, ok := comments[pathKey(valPath)]; ok {
				enumVal.DocComment = cleanComment(loc.GetLeadingComments())
			}
		}

		m.Options = append(m.Options, enumVal)
	}

	return m
}
