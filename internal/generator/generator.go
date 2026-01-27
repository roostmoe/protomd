package generator

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/roostmoe/protomd/internal/mdgen"
	"github.com/roostmoe/protomd/internal/registry"
	"github.com/roostmoe/protomd/internal/walk"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// Generator handles the file generation logic
type Generator struct {
	Registry       *registry.Registry
	Request        *pluginpb.CodeGeneratorRequest
	Response       *pluginpb.CodeGeneratorResponse
	SkipNamespaces int // number of namespace segments to skip from the beginning
}

// New creates a new Generator from a CodeGeneratorRequest
func New(req *pluginpb.CodeGeneratorRequest) *Generator {
	g := &Generator{
		Registry: registry.Build(req),
		Request:  req,
		Response: &pluginpb.CodeGeneratorResponse{
			SupportedFeatures: proto.Uint64(uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)),
		},
		SkipNamespaces: 0,
	}

	// parse parameters from request
	g.parseParameters(req.GetParameter())

	return g
}

// parseParameters extracts options from the protoc parameter string
func (g *Generator) parseParameters(params string) {
	if params == "" {
		return
	}

	// parameters are comma-separated key=value pairs
	for param := range strings.SplitSeq(params, ",") {
		kv := strings.SplitN(param, "=", 2)
		if len(kv) != 2 {
			continue
		}

		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "skip-namespaces":
			if n, err := strconv.Atoi(value); err == nil && n >= 0 {
				g.SkipNamespaces = n
			}
		}
	}
}

// Generate walks through all files and generates output
func (g *Generator) Generate() *pluginpb.CodeGeneratorResponse {
	// group files by namespace (package)
	namespaces := g.groupByNamespace()

	// walk each namespace and generate files
	for pkg, files := range namespaces {
		g.generateNamespace(pkg, files)
	}

	return g.Response
}

// groupByNamespace organizes files by their package
func (g *Generator) groupByNamespace() map[string][]*descriptorpb.FileDescriptorProto {
	namespaces := make(map[string][]*descriptorpb.FileDescriptorProto)
	for _, fname := range g.Request.GetFileToGenerate() {
		f := g.Registry.FilesByName[fname]
		pkg := f.GetPackage()
		namespaces[pkg] = append(namespaces[pkg], f)
	}
	return namespaces
}

// generateNamespace handles file generation for a single namespace
func (g *Generator) generateNamespace(pkg string, files []*descriptorpb.FileDescriptorProto) {
	// transform the package name based on skip-namespaces setting
	outputPath := g.transformNamespace(pkg)

	// collect services and messages with their metadata
	g.collectAndGenerate(pkg, outputPath, files)
}

// transformNamespace applies the skip-namespaces setting to a package name
// e.g. with skip-namespaces=1, "infrascope.admin.v1" becomes "admin/v1"
func (g *Generator) transformNamespace(pkg string) string {
	if g.SkipNamespaces == 0 {
		return strings.ReplaceAll(pkg, ".", "/")
	}

	parts := strings.Split(pkg, ".")
	if g.SkipNamespaces >= len(parts) {
		// if we're skipping more segments than exist, return the last segment
		return parts[len(parts)-1]
	}

	// skip the first N segments
	remaining := parts[g.SkipNamespaces:]
	return strings.Join(remaining, "/")
}

// symbolMetadata tracks the file and path for a symbol
type symbolMetadata struct {
	fileName string
	path     []int32
	parent   string // parent name for nested symbols (e.g., "Message" for nested enum "Message.Status")
	proto    any    // *descriptorpb.ServiceDescriptorProto or *descriptorpb.DescriptorProto or *descriptorpb.EnumDescriptorProto
}

// collectAndGenerate walks files and generates content using mdgen
func (g *Generator) collectAndGenerate(originalPkg, outputPath string, files []*descriptorpb.FileDescriptorProto) {
	services := make(map[string]symbolMetadata)
	messages := make(map[string]symbolMetadata)
	enums := make(map[string]symbolMetadata)

	// walk files and collect symbols with their paths
	for _, f := range files {
		fileName := f.GetName()

		walk.WalkFile(f, walk.Visitor{
			OnService: func(path []int32, s *descriptorpb.ServiceDescriptorProto) {
				services[s.GetName()] = symbolMetadata{
					fileName: fileName,
					path:     path,
					parent:   "",
					proto:    s,
				}
			},
			OnMessage: func(path []int32, m *descriptorpb.DescriptorProto) {
				// collect top-level messages
				messages[m.GetName()] = symbolMetadata{
					fileName: fileName,
					path:     path,
					parent:   "",
					proto:    m,
				}
				// recursively collect nested messages and enums
				g.collectNested(fileName, path, m.GetName(), messages, enums, m)
			},
			OnEnum: func(path []int32, e *descriptorpb.EnumDescriptorProto) {
				// collect top-level enums
				enums[e.GetName()] = symbolMetadata{
					fileName: fileName,
					path:     path,
					parent:   "",
					proto:    e,
				}
			},
		})
	}

	// collect all enums for namespace index
	var mdEnums []mdgen.Enum
	for _, meta := range enums {
		enum := meta.proto.(*descriptorpb.EnumDescriptorProto)
		mdEnum := mdgen.NewEnum(g.Registry, enum, meta.fileName, meta.path)
		// apply parent prefix if nested
		if meta.parent != "" {
			mdEnum.Name = meta.parent + "." + mdEnum.Name
		}
		mdEnums = append(mdEnums, mdEnum)
	}

	// collect all messages for namespace index
	var mdMessages []mdgen.Message
	for _, meta := range messages {
		msg := meta.proto.(*descriptorpb.DescriptorProto)
		mdMsg := mdgen.NewMessage(g.Registry, msg, meta.fileName, meta.path)
		// apply parent prefix if nested
		if meta.parent != "" {
			mdMsg.Name = meta.parent + "." + mdMsg.Name
		}
		mdMessages = append(mdMessages, mdMsg)
	}

	// collect all services for namespace index
	var mdServices []mdgen.Service
	for _, meta := range services {
		svc := meta.proto.(*descriptorpb.ServiceDescriptorProto)
		mdSvc := mdgen.NewService(g.Registry, svc, meta.fileName, meta.path)
		mdServices = append(mdServices, mdSvc)
	}

	// generate namespace index with all messages, enums and services
	nsIndexPath := fmt.Sprintf("/%s/index.md", outputPath)
	namespace := mdgen.NewNamespace(originalPkg, mdServices, mdEnums, mdMessages)
	var nsBuf bytes.Buffer
	if err := namespace.Write(&nsBuf); err != nil {
		panic(fmt.Sprintf("failed to render namespace %s: %v", originalPkg, err))
	}
	g.addFile(nsIndexPath, nsBuf.String())

	// generate service files using mdgen
	for svcName, meta := range services {
		svc := meta.proto.(*descriptorpb.ServiceDescriptorProto)

		// create mdgen service
		mdSvc := mdgen.NewService(g.Registry, svc, meta.fileName, meta.path)

		// generate grpc service index
		var buf bytes.Buffer
		if err := mdSvc.WriteGrpc(&buf); err != nil {
			panic(fmt.Sprintf("failed to render grpc service %s: %v", svcName, err))
		}
		svcIndexPath := fmt.Sprintf("/%s/rpc/%s/index.md", outputPath, svcName)
		g.addFile(svcIndexPath, buf.String())

		// generate rest service index if needed
		if mdSvc.HasHttp {
			var restBuf bytes.Buffer
			if err := mdSvc.WriteRest(&restBuf); err != nil {
				panic(fmt.Sprintf("failed to render rest service %s: %v", svcName, err))
			}
			restIndexPath := fmt.Sprintf("/%s/rest/%s/index.md", outputPath, svcName)
			g.addFile(restIndexPath, restBuf.String())
		}

		// generate method files
		for i, method := range svc.GetMethod() {
			methodPath := append(append([]int32{}, meta.path...), 2, int32(i))
			mdMethod := mdgen.NewMethod(g.Registry, method, meta.fileName, methodPath)

			// generate grpc method
			var methodBuf bytes.Buffer
			if err := mdMethod.WriteGrpc(&methodBuf); err != nil {
				panic(fmt.Sprintf("failed to render grpc method %s.%s: %v", svcName, method.GetName(), err))
			}
			methodFilePath := fmt.Sprintf("/%s/rpc/%s/%s.md", outputPath, svcName, method.GetName())
			g.addFile(methodFilePath, methodBuf.String())

			// generate rest method if needed
			if mdMethod.HasHttp {
				var restMethodBuf bytes.Buffer
				if err := mdMethod.WriteRest(&restMethodBuf); err != nil {
					panic(fmt.Sprintf("failed to render rest method %s.%s: %v", svcName, method.GetName(), err))
				}
				restMethodPath := fmt.Sprintf("/%s/rest/%s/%s.md", outputPath, svcName, method.GetName())
				g.addFile(restMethodPath, restMethodBuf.String())
			}
		}
	}
}

// collectNested recursively collects nested messages and enums from a message
func (g *Generator) collectNested(
	fileName string,
	parentPath []int32,
	parentName string,
	messages map[string]symbolMetadata,
	enums map[string]symbolMetadata,
	msg *descriptorpb.DescriptorProto,
) {
	// collect nested enums
	for i, enum := range msg.GetEnumType() {
		// path: parentPath + [4 (enum_type field), i]
		enumPath := append(append([]int32{}, parentPath...), 4, int32(i))
		fullName := parentName + "." + enum.GetName()

		enums[fullName] = symbolMetadata{
			fileName: fileName,
			path:     enumPath,
			parent:   parentName,
			proto:    enum,
		}
	}

	// collect nested messages
	for i, nested := range msg.GetNestedType() {
		// path: parentPath + [3 (nested_type field), i]
		nestedPath := append(append([]int32{}, parentPath...), 3, int32(i))
		fullName := parentName + "." + nested.GetName()

		messages[fullName] = symbolMetadata{
			fileName: fileName,
			path:     nestedPath,
			parent:   parentName,
			proto:    nested,
		}

		// recursively collect from this nested message
		g.collectNested(fileName, nestedPath, fullName, messages, enums, nested)
	}
}

// addFile creates a CodeGeneratorResponse_File and appends it to the response
//
// To write actual content:
// 1. Build your markdown/content as a string
// 2. Pass it to this function
// 3. protoc will write file.Name with file.Content to disk
//
// Example usage:
//
//	content := buildMarkdownContent(data)
//	g.addFile("/path/to/output.md", content)
func (g *Generator) addFile(path, content string) {
	// protoc wants relative paths (no leading slash)
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	file := &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(path),
		Content: proto.String(content),
	}
	g.Response.File = append(g.Response.File, file)
}
