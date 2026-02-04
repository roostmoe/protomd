package registry

import (
	"fmt"
	"strings"

	"github.com/roostmoe/protomd/internal/models"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Registry struct {
	FilesByName map[string]*descriptorpb.FileDescriptorProto

	// Packages indexed by package path
	Packages map[string]models.Package

	// comments indexed by source path, per file
	// file name -> ("6,0,2,1" -> location)
	Comments map[string]models.CommentMap
}

func Build(req *pluginpb.CodeGeneratorRequest) *Registry {
	r := &Registry{
		FilesByName:   map[string]*descriptorpb.FileDescriptorProto{},
		Packages:      map[string]models.Package{},
		Comments:      map[string]models.CommentMap{},
	}

	for _, f := range req.GetProtoFile() {
		name := f.GetName()
		pkg := f.GetPackage()

		pkgModel := &models.Package{
			Name: pkg,
			Messages: []models.Message{},
			Services: []models.Service{},
			Enums: 		[]models.Enum{},
		}

		if pkg, ok := r.Packages[pkg]; ok {
			pkgModel = &pkg
		}

		r.FilesByName[name] = f
		r.Comments[name] = indexComments(f)

		for idx, m := range f.GetMessageType() {
			path := []int32{4, int32(idx)}
			pkgModel.AddProtoMessage(m, r.Comments[name], path)
		}

		for idx, e := range f.GetEnumType() {
			path := []int32{4, int32(idx)}
			pkgModel.AddProtoEnum(e, r.Comments[name], path)
		}

		for idx, s := range f.GetService() {
			path := []int32{4, int32(idx)}
			pkgModel.AddProtoService(s, r.Comments[name], path)
		}

		r.Packages[pkg] = *pkgModel
	}

	return r
}

// ---- comments

func indexComments(f *descriptorpb.FileDescriptorProto) models.CommentMap {
	out := models.CommentMap{}
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
		fmt.Fprint(&b, p)
	}
	return b.String()
}
