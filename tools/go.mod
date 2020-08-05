module github.com/percona-platform/dbaas-api/tools

go 1.14

// some dependecies should be synced with Dockerfile

require (
	github.com/golang/protobuf v1.4.2
	github.com/golangci/golangci-lint v1.30.0
	github.com/mwitkow/go-proto-validators v0.3.0
	github.com/quasilyte/go-consistent v0.0.0-20200404105227-766526bf1e96
	github.com/reviewdog/reviewdog v0.10.0
	github.com/shurcooL/githubv4 v0.0.0-20200414012201-bbc966b061dd
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f // indirect
	github.com/uber/prototool v1.10.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/tools v0.0.0-20200724022722-7017fd6b1305
	google.golang.org/grpc v1.30.0
)
