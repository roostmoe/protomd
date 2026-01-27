package mdgen

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

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

// pathKey converts a source path to a lookup key for the comments map
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

// cleanComment trims whitespace and normalizes comment text
func cleanComment(s string) string {
	return strings.TrimSpace(s)
}

func tableEscape(s string) string {
	ns := s
	ns = strings.ReplaceAll(ns, "\n\n", "<br>")
	ns = strings.ReplaceAll(ns, "\n", "&nbsp;")
	return ns
}
