// +build tools

package tools // import "github.com/percona-platform/dbaas-api/tools"

import (
	// code generators plus their dependencies (to make them direct in go.mod)
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/mwitkow/go-proto-validators/protoc-gen-govalidators"
	_ "github.com/uber/prototool/cmd/prototool"
	_ "google.golang.org/grpc"

	// other tools
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/quasilyte/go-consistent"
	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
	_ "golang.org/x/tools/cmd/goimports"
)
