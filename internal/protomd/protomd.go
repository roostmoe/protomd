package protomd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/roostmoe/protomd/internal/mdgen"
	"github.com/roostmoe/protomd/internal/registry"
	"google.golang.org/protobuf/types/pluginpb"
)

type Config struct {
	SkipNamespaces int
}

func ReadConfig(params string) *Config {
	config := &Config{}

	if params == "" {
		return config
	}

	for param := range strings.SplitSeq(params, ",") {
		kv := strings.SplitN(param, "=", 2)
		if len(kv) != 2 {
			continue
		}

		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "skip_namespaces":
			if n, err := strconv.Atoi(value); err == nil && n >= 0 {
				config.SkipNamespaces = n
			}
		}
	}

	return config
}

type Protomd struct {
	Cfg      *Config
	Registry *registry.Registry
	req      *pluginpb.CodeGeneratorRequest
}

func New(req *pluginpb.CodeGeneratorRequest) *Protomd {
	return &Protomd{
		Cfg:      ReadConfig(req.GetParameter()),
		Registry: registry.Build(req),
		req:      req,
	}
}

func (p *Protomd) Run() (*pluginpb.CodeGeneratorResponse, error) {
	resp := &pluginpb.CodeGeneratorResponse{}

	for _, pkg := range p.Registry.Packages {
		fmt.Fprintf(os.Stderr, "processing package %v:\n", pkg.Name)
		fileName := fmt.Sprintf("%v/index.md", pkg.Name)

		md, err := mdgen.Render(pkg)
		if err != nil {
			return nil, fmt.Errorf("failed to build package %v: %v", pkg.Name, err)
		}

		resp.File = append(resp.File, &pluginpb.CodeGeneratorResponse_File{
			Name: &fileName,
			Content: &md,
		})

		for _, msg := range pkg.Messages {
			fmt.Fprintf(os.Stderr, "  processing message %v:\n", pkg.Name)

			for _, fld := range msg.Fields {
				fmt.Fprintf(os.Stderr, "    processing field %v:\n", fld.Name)
			}

			// for _, enm := range msg.Enums {
			// 	fmt.Fprintf(os.Stderr, "    processing enum %v:", enm.Name)
			// }
		}

		for _, enm := range pkg.Enums {
			fmt.Fprintf(os.Stderr, "  processing enum %v:\n", enm.Name)

			for _, env := range enm.Values {
				fmt.Fprintf(os.Stderr, "    processing enum value %v:\n", env.Name)
			}
		}

		for _, svc := range pkg.Services {
			fmt.Fprintf(os.Stderr, "  processing service %v:\n", svc.Name)
		}
	}

	return resp, nil
}
