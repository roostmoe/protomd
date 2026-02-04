package models

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type CommentMap map[string]*descriptorpb.SourceCodeInfo_Location

func getExtension[A any](opts proto.Message, e protoreflect.ExtensionType) (*A, bool) {
	if !proto.HasExtension(opts, e) {
		return nil, false
	}

	ext := proto.GetExtension(opts, e)
	if ext == nil {
		return nil, false
	}

	val, ok := ext.(*A)
	if !ok {
		return nil, false
	}

	return val, true
}

func getMethodExtension[A any](opts *descriptorpb.MethodOptions, e protoreflect.ExtensionType) (*A, bool) {
	if opts == nil {
		return nil, false
	}
	return getExtension[A](opts, e)
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

func getComment(comments CommentMap, path []int32) string {
	loc, ok := comments[pathKey(path)]
	if !ok {
		return ""
	}
	return cleanComment(loc.GetLeadingComments())
}

func cleanComment(s string) string {
	return strings.TrimSpace(s)
}
