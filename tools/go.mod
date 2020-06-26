module github.com/percona-platform/dbaas-api/tools

go 1.14

// some dependecies should be synced with Dockerfile

require (
	github.com/golang/protobuf v1.3.5
	github.com/golangci/golangci-lint v1.27.0
	github.com/mwitkow/go-proto-validators v0.3.0
	github.com/quasilyte/go-consistent v0.0.0-20200404105227-766526bf1e96
	github.com/reviewdog/reviewdog v0.10.0
	github.com/uber/prototool v1.10.0
	golang.org/x/tools v0.0.0-20200623045635-ff88973b1e4e
	google.golang.org/grpc v1.30.0
)
