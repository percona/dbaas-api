module github.com/percona-platform/dbaas-api

go 1.17

// some dependecies should be synced with Dockerfile

require (
	github.com/golang/protobuf v1.5.2
	github.com/mwitkow/go-proto-validators v0.3.2
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/gogo/protobuf v1.3.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)
