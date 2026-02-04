package main

import (
	"fmt"
	"io"
	"os"

	"github.com/roostmoe/protomd/internal/protomd"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	req, err := readRequest()
	if err != nil {
		fatal("failed to read request: %v", err)
	}

	p := protomd.New(req)

	resp, err := p.Run()
	if err != nil {
		fatal("failed to build plugin response: %v", err)
	}

	if err := writeResponse(resp); err != nil {
		fatal("failed to write response: %v", err)
	}
}

// readRequest reads and unmarshals the CodeGeneratorRequest from stdin
func readRequest() (*pluginpb.CodeGeneratorRequest, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("read stdin: %w", err)
	}

	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(data, req); err != nil {
		return nil, fmt.Errorf("unmarshal request: %w", err)
	}

	return req, nil
}

// writeResponse marshals and writes the CodeGeneratorResponse to stdout
func writeResponse(resp *pluginpb.CodeGeneratorResponse) error {
	data, err := proto.Marshal(resp)
	if err != nil {
		return fmt.Errorf("marshal response: %w", err)
	}

	if _, err := os.Stdout.Write(data); err != nil {
		return fmt.Errorf("write stdout: %w", err)
	}

	return nil
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "protoc-gen-md: "+format+"\n", args...)
	os.Exit(1)
}
