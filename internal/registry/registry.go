package registry

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Registry struct {
	FilesByName map[string]*descriptorpb.FileDescriptorProto

	// fully-qualified proto names (leading dot), e.g. ".foo.v1.Bar"
	Messages map[string]*descriptorpb.DescriptorProto
	Enums    map[string]*descriptorpb.EnumDescriptorProto
	Services map[string]*descriptorpb.ServiceDescriptorProto

	// method index is optional but handy for crosslinks / lookups
	// key: ".pkg.Service/Method"
	Methods map[string]*descriptorpb.MethodDescriptorProto

	// comments indexed by source path, per file
	// file name -> ("6,0,2,1" -> location)
	Comments map[string]map[string]*descriptorpb.SourceCodeInfo_Location

	// also useful for linking symbols back to files/packages
	PackageByFile map[string]string
}

func Build(req *pluginpb.CodeGeneratorRequest) *Registry {
	r := &Registry{
		FilesByName:   map[string]*descriptorpb.FileDescriptorProto{},
		Messages:      map[string]*descriptorpb.DescriptorProto{},
		Enums:         map[string]*descriptorpb.EnumDescriptorProto{},
		Services:      map[string]*descriptorpb.ServiceDescriptorProto{},
		Methods:       map[string]*descriptorpb.MethodDescriptorProto{},
		Comments:      map[string]map[string]*descriptorpb.SourceCodeInfo_Location{},
		PackageByFile: map[string]string{},
	}

	for _, f := range req.GetProtoFile() {
		name := f.GetName()
		pkg := f.GetPackage()

		r.FilesByName[name] = f
		r.PackageByFile[name] = pkg
		r.Comments[name] = indexComments(f)

		// top-level + nested messages/enums
		for _, m := range f.GetMessageType() {
			indexMessage(r, pkg, "", m)
		}
		for _, e := range f.GetEnumType() {
			full := fq(pkg, "", e.GetName())
			r.Enums[full] = e
		}

		// services + methods
		for _, s := range f.GetService() {
			svcFull := fq(pkg, "", s.GetName())
			r.Services[svcFull] = s

			for _, m := range s.GetMethod() {
				key := svcFull + "/" + m.GetName()
				r.Methods[key] = m
			}
		}
	}

	return r
}

func indexMessage(r *Registry, pkg, parent string, m *descriptorpb.DescriptorProto) {
	msgFull := fq(pkg, parent, m.GetName())
	r.Messages[msgFull] = m

	// nested enums inside this message
	nextParent := m.GetName()
	if parent != "" {
		nextParent = parent + "." + m.GetName()
	}

	for _, e := range m.GetEnumType() {
		eFull := fq(pkg, nextParent, e.GetName())
		r.Enums[eFull] = e
	}

	// nested messages (recursively)
	for _, nested := range m.GetNestedType() {
		indexMessage(r, pkg, nextParent, nested)
	}
}

// fq builds fully-qualified proto names with the leading dot.
// parent is the containing message chain like "Outer.Inner".
func fq(pkg, parent, name string) string {
	// ".<pkg>.<parent>.<name>" (skipping empties)
	full := "." + pkg
	if parent != "" {
		full += "." + parent
	}
	full += "." + name
	return full
}

// ---- comments

func indexComments(f *descriptorpb.FileDescriptorProto) map[string]*descriptorpb.SourceCodeInfo_Location {
	out := map[string]*descriptorpb.SourceCodeInfo_Location{}
	sci := f.GetSourceCodeInfo()
	if sci == nil {
		return out
	}
	for _, loc := range sci.GetLocation() {
		out[pathKey(loc.GetPath())] = loc
	}
	return out
}

func pathKey(path []int32) string {
	if len(path) == 0 {
		return ""
	}
	var b strings.Builder
	for i, p := range path {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(fmt.Sprint(p))
	}
	return b.String()
}
