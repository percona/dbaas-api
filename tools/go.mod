module github.com/percona-platform/dbaas-api/tools

go 1.14

// some dependecies should be synced with Dockerfile

require (
	github.com/fullstorydev/grpcurl v1.7.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/golangci/golangci-lint v1.31.0
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/quasilyte/go-consistent v0.0.0-20200404105227-766526bf1e96
	github.com/reviewdog/reviewdog v0.10.2
	github.com/uber/prototool v1.10.0
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/tools v0.0.0-20200908211811-12e1bf57a112
	google.golang.org/grpc v1.32.0
)
