package mdgen

import (
	"fmt"
	"io"

	md "github.com/nao1215/markdown"
	"github.com/roostmoe/protomd/internal/registry"
	"github.com/roostmoe/protomd/internal/wellknown"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Message struct {
	Name         string
	DocComment   string
	Fields       []Field
	Oneofs       []Oneof
	IsWellKnown  bool // true if this is a well-known type
	WellKnownUrl string
	TypeName     string // fully-qualified type name (e.g., ".google.protobuf.Timestamp")
}

type Oneof struct {
	Name       string
	DocComment string
	Fields     []Field
}

func NewMessage(
	reg *registry.Registry,
	proto *descriptorpb.DescriptorProto,
	fileName string,
	path []int32,
) Message {
	m := Message{
		Name:   proto.GetName(),
		Fields: []Field{},
		Oneofs: []Oneof{},
	}

	// construct the fully-qualified type name
	if pkg, ok := reg.PackageByFile[fileName]; ok {
		if pkg != "" {
			m.TypeName = fmt.Sprintf(".%s.%s", pkg, proto.GetName())
		} else {
			m.TypeName = fmt.Sprintf(".%s", proto.GetName())
		}

		// check if this is a well-known type
		m.IsWellKnown = wellknown.IsWellKnown(m.TypeName)
		if m.IsWellKnown {
			m.WellKnownUrl = wellknown.GetDocURL(m.TypeName)
		}
	}

	// get comment from registry
	if comments, ok := reg.Comments[fileName]; ok {
		if loc, ok := comments[pathKey(path)]; ok {
			m.DocComment = cleanComment(loc.GetLeadingComments())
		}
	}

	// build all fields first
	allFields := []Field{}
	for i, f := range proto.GetField() {
		fieldPath := append(append([]int32{}, path...), 2, int32(i)) // 2 = field tag in DescriptorProto
		allFields = append(allFields, NewField(reg, f, fileName, fieldPath))
	}

	// extract oneof declarations if present
	if len(proto.GetOneofDecl()) > 0 {
		// create oneof structures
		for i, oneofDecl := range proto.GetOneofDecl() {
			oneof := Oneof{
				Name:   oneofDecl.GetName(),
				Fields: []Field{},
			}

			// get comment for this oneof
			if comments, ok := reg.Comments[fileName]; ok {
				oneofPath := append(append([]int32{}, path...), 8, int32(i)) // 8 = oneof_decl tag
				if loc, ok := comments[pathKey(oneofPath)]; ok {
					oneof.DocComment = cleanComment(loc.GetLeadingComments())
				}
			}

			// collect fields that belong to this oneof
			for _, field := range allFields {
				if field.OneofIndex != nil && int(*field.OneofIndex) == i {
					oneof.Fields = append(oneof.Fields, field)
				}
			}

			m.Oneofs = append(m.Oneofs, oneof)
		}

		// regular fields are those not in any oneof
		for _, field := range allFields {
			if field.OneofIndex == nil {
				m.Fields = append(m.Fields, field)
			}
		}
	} else {
		// no oneofs, all fields are regular
		m.Fields = allFields
	}

	return m
}

func (m Message) Build(b *md.Markdown) *md.Markdown {
	var fieldRows [][]string
	for _, f := range m.Fields {
		fieldRows = append(fieldRows, f.BuildTableRow())
	}

	mdComment := tableEscape(m.DocComment)

	return b.H2f("%v\n", md.Code(m.Name)).
		PlainTextf("%v\n", mdComment).
		Table(md.TableSet{
			Header: []string{"Fields", ""},
			Rows:   fieldRows,
		})
}

func (m Message) Write(w io.Writer) error {
	type data struct {
		Message Message
	}

	return tmpl.ExecuteTemplate(w, "message.md", data{Message: m})
}
