# protomd Agent Guide

protoc plugin that generates markdown documentation from protobuf files

## Build/Test Commands

### Build
```bash
just build                    # builds to bin/protoc-gen-md
```

### Test
```bash
go test ./...                 # run all tests
go test ./internal/registry   # run tests for specific package
go test -v ./...              # verbose output
go test -run TestName ./...   # run single test by name
go test -cover ./...          # with coverage
```

### Linting
```bash
go vet ./...                  # standard go vet
gofmt -s -w .                 # format code (use -s for simplification)
```

### Dependencies
```bash
go mod tidy                   # clean up dependencies
go mod download               # download dependencies
```

## Project Structure

```
cmd/protoc-gen-md/    - main entry point for protoc plugin
internal/
  generator/          - core generation logic, namespace handling
  mdgen/              - markdown generation (services, messages, fields, enums)
  registry/           - protobuf descriptor registry, comment indexing
  walk/               - file descriptor traversal utilities
  wellknown/          - well-known type documentation URLs
```

## Code Style

### Imports
- Standard library first
- Third-party packages second
- Local packages last
- Groups separated by blank lines
- No named imports unless necessary for disambiguation

Example:
```go
import (
    "fmt"
    "strings"

    "github.com/Masterminds/sprig/v3"
    "google.golang.org/protobuf/types/descriptorpb"

    "github.com/roostmoe/protomd/internal/registry"
)
```

### Naming
- Use short, descriptive variable names: `m` for message, `f` for field, `s` for service
- Exported types/functions start with capital letter
- Unexported helpers in lowercase
- Acronyms stay uppercase: `URL`, `HTTP`, `RPC` (not `Url`, `Http`, `Rpc`)
- Registry keys use fully-qualified names with leading dot: `.pkg.Message`
- Source code paths use int32 slices matching descriptor.proto field numbers

### Types
- Prefer `*descriptorpb.DescriptorProto` over `interface{}`
- Use pointer receivers for methods that modify state
- Document exported types with godoc comments
- Use type aliases sparingly

### Error Handling
- Use `fmt.Errorf` with `%w` for wrapping errors
- Use `panic()` for unrecoverable plugin errors (protoc expects error status via stderr)
- Check errors immediately after calls
- No naked returns in error paths

Example:
```go
if err := doSomething(); err != nil {
    return fmt.Errorf("do something: %w", err)
}
```

### Comments
- Godoc format for exported symbols
- Leading comments describe **why** not what
- Use `//` for single-line comments
- Inline phase comments for complex multi-step logic (see main.go)

### Functions
- Keep functions focused and small
- Exported functions get godoc comments
- Helper functions grouped logically
- Prefer explicit parameters over global state

### Maps and Slices
- Use `make()` with capacity hints for large collections
- Append idiom: `slice = append(slice, item)`
- Clone slices before mutation: `append([]int32{}, path...)`
- Map iteration order is undefined - sort keys if order matters

### Protobuf Patterns
- Source code paths are `[]int32` indexes into descriptor tree
- Path format: `[field_number, index, field_number, index, ...]`
- Comments keyed by comma-joined path string: `"6,0,2,1"`
- Get descriptors with `proto.GetName()`, `proto.GetPackage()`, etc.
- Check extensions with `proto.HasExtension()` before `proto.GetExtension()`

### Templates
- Templates live in `internal/mdgen/templates/`
- Use sprig functions for string manipulation
- Custom template functions in mdgen.go init()
- Execute with `tmpl.ExecuteTemplate(w, "template.md", data)`

### File Generation
- Use `generator.addFile(path, content)` to add output files
- Paths must be relative (no leading `/`)
- Content is raw string written by protoc
- One namespace = one index.md + service/method .md files

### Registry Usage
- Registry built once from CodeGeneratorRequest
- Messages/Enums/Services indexed by fully-qualified name
- Comments indexed per-file by source path
- Use registry for cross-references and type lookups

### Testing
- No tests currently in codebase
- When adding tests: use `_test.go` suffix
- Use table-driven tests for multiple cases
- Mock registry for unit tests of mdgen components

## Common Tasks

### Adding a new field type
1. Update `Field` struct in `internal/mdgen/field.go`
2. Extend `NewField` constructor logic
3. Update field template if needed

### Adding template functions
1. Add function to `internal/mdgen/mdgen.go`
2. Register in `init()` funcMap
3. Use in template with `{{ functionName args }}`

### Handling new annotations
1. Import extension from `google.golang.org/genproto`
2. Use `getExtension()` or `getMethodExtension()` helpers
3. Add to relevant mdgen struct (Service, Method, Field, etc.)

### Debugging generated output
1. Build: `just build`
2. Run protoc with plugin on test .proto files
3. Inspect generated .md files
4. Check source paths match expected descriptor tree structure

## Go Version
- Requires: **1.25.6** (see mise.toml and go.mod)
- Use mise or install manually

## Dependencies
- google.golang.org/protobuf - protobuf descriptor types
- google.golang.org/genproto - googleapis annotations
- github.com/Masterminds/sprig/v3 - template functions
