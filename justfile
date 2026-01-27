build:
  go build -o bin/protoc-gen-md ./cmd/protoc-gen-md

example: build
  cd example && buf generate
