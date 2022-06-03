//go:build tools
// +build tools

package tools

import (
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/ramya-rao-a/go-outline"
	_ "github.com/uudashr/gopkgs/v2/cmd/gopkgs"
	_ "golang.org/x/tools/gopls"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
