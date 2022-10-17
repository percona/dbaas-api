module github.com/percona-platform/dbaas-api

go 1.17

// some dependecies should be synced with Dockerfile

require (
	github.com/golang/protobuf v1.5.2
	github.com/mwitkow/go-proto-validators v0.3.2
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/gogo/protobuf v1.3.0 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
