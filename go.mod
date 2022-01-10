module github.com/percona-platform/dbaas-api

go 1.17

// some dependecies should be synced with Dockerfile

require (
	github.com/golang/protobuf v1.5.0
	github.com/mwitkow/go-proto-validators v0.3.2
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.26.0
)

require (
	github.com/gogo/protobuf v1.3.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
