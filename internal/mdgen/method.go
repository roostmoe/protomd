package mdgen

import (
	"fmt"
	"io"
	"os"
	"strings"

	md "github.com/nao1215/markdown"
	"github.com/roostmoe/protomd/internal/registry"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Method struct {
	Name       string
	DocComment string

	RestPath   string
	RestMethod string

	HasHttp bool

	Input  Message
	Output Message
}

func NewMethod(
	reg *registry.Registry,
	proto *descriptorpb.MethodDescriptorProto,
	fileName string,
	path []int32,
) Method {
	m := Method{
		Name: proto.GetName(),
	}

	// get comment from registry
	if comments, ok := reg.Comments[fileName]; ok {
		if loc, ok := comments[pathKey(path)]; ok {
			m.DocComment = cleanComment(loc.GetLeadingComments())
		}
	}

	// resolve input/output messages from registry
	inputType := proto.GetInputType()
	if inputMsg, ok := reg.Messages[inputType]; ok {
		// we need to find the file for this message to get proper paths
		// for now, create message without path (comments will be empty)
		m.Input = NewMessage(reg, inputMsg, "", nil)
	}

	outputType := proto.GetOutputType()
	if outputMsg, ok := reg.Messages[outputType]; ok {
		m.Output = NewMessage(reg, outputMsg, "", nil)
	}

	// check for HTTP annotations
	m.HasHttp = m.checkHttp(proto)

	return m
}

func (m Method) BuildSummary(b *md.Markdown) *md.Markdown {
	b = b.H2(m.Name).PlainText(m.DocComment).PlainText("\n")

	inputUrl := strings.ToLower(fmt.Sprintf("../#%s", m.Input.Name))
	if m.Input.IsWellKnown {
		inputUrl = m.Input.WellKnownUrl
	}

	outputUrl := strings.ToLower(fmt.Sprintf("../#%s", m.Output.Name))
	if m.Output.IsWellKnown {
		outputUrl = m.Output.WellKnownUrl
	}

	rpcLine := fmt.Sprintf(
		`<pre><code>rpc %s (<a href="%s">%s</a>) returns (<a href="%s">%s</a>)</code></pre>`,
		m.Name,
		inputUrl,
		m.Input.Name,
		outputUrl,
		m.Output.Name,
	)

	b = b.PlainTextf(`=== "gRPC"
    %s`, rpcLine)

	if m.RestMethod != "" && m.RestPath != "" {
		b = b.PlainText(`=== "HTTP"`)

		if m.RestMethod == "GET" {
			b = m.buildHttpROSummary(b)
		} else {
			b = m.buildHttpRWSummary(b)
		}
	}

	return b
}

func (m *Method) buildHttpROSummary(b *md.Markdown) *md.Markdown {
	cb := fmt.Sprintf(
		"<pre><code>%s</code></pre>",
		fmt.Sprintf(`%s %s`, m.RestMethod, m.RestPath),
	)

	newCb := ""
	for line := range strings.SplitSeq(cb, "\n") {
		newCb = newCb + fmt.Sprintf("    %s\n", line)
	}

	return b.PlainText(newCb)
}

func (m *Method) buildHttpRWSummary(b *md.Markdown) *md.Markdown {
	cb := fmt.Sprintf(
		"<pre><code>%s</code></pre>",
		fmt.Sprintf(`%s %s`, m.RestMethod, m.RestPath),
	)

	newCb := ""
	for line := range strings.SplitSeq(cb, "\n") {
		newCb = newCb + fmt.Sprintf("    %s\n", line)
	}

	return b.PlainText(newCb)
}

func (m *Method) checkHttp(proto *descriptorpb.MethodDescriptorProto) bool {
	opts := proto.GetOptions()
	if opts == nil {
		fmt.Fprintf(os.Stderr, "method %s has no options\n", m.Name)
		return false
	}

	rule, ok := getMethodExtension[annotations.HttpRule](opts, annotations.E_Http)
	if !ok {
		fmt.Fprintf(os.Stderr, "method %s has no http annotation\n", m.Name)
		return false
	}

	switch r := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		m.RestMethod = "GET"
		m.RestPath = r.Get
	case *annotations.HttpRule_Post:
		m.RestMethod = "POST"
		m.RestPath = r.Post
	case *annotations.HttpRule_Patch:
		m.RestMethod = "PATCH"
		m.RestPath = r.Patch
	case *annotations.HttpRule_Put:
		m.RestMethod = "PUT"
		m.RestPath = r.Put
	case *annotations.HttpRule_Delete:
		m.RestMethod = "DELETE"
		m.RestPath = r.Delete
	case *annotations.HttpRule_Custom:
		m.RestMethod = r.Custom.Kind
		m.RestPath = r.Custom.Path
	default:
		m.RestPath = "unknown"
		return false
	}

	fmt.Fprintf(os.Stderr, "method %s has method %+v and path %+v\n", m.Name, m.RestMethod, m.RestPath)

	return true
}

func (m Method) WriteGrpc(w io.Writer) error {
	type data struct {
		Method Method
	}

	return tmpl.ExecuteTemplate(w, "grpc-method.md", data{Method: m})
}

func (m Method) WriteRest(w io.Writer) error {
	type data struct {
		Method Method
	}

	return tmpl.ExecuteTemplate(w, "rest-method.md", data{Method: m})
}
