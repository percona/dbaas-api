DOCKER_DEV_IMAGE  = percona-platform-dbaas-prototool:dev
DOCKER_RUN_IMAGE ?= docker.pkg.github.com/percona-platform/dbaas-api/dbaas-prototool:latest
DOCKER_RUN_CMD    = docker run --rm --mount='type=bind,src=$(PWD),dst=/work' $(DOCKER_RUN_IMAGE)

default: help

help:                                      ## Display this help message
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

init:                                      ## Install development tools
	go build -modfile=tools/go.mod -o bin/goimports golang.org/x/tools/cmd/goimports
	go build -modfile=tools/go.mod -o bin/stringer golang.org/x/tools/cmd/stringer
	go build -modfile=tools/go.mod -o bin/go-consistent github.com/quasilyte/go-consistent
	go build -modfile=tools/go.mod -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
	go build -modfile=tools/go.mod -o bin/reviewdog github.com/reviewdog/reviewdog/cmd/reviewdog

ci-init:                ## Initialize CI environment
	# nothing there yet

gen:                                       ## Format, check, and generate using prototool Docker image
	# $(DOCKER_RUN_CMD) prototool break check api/controller -f api/controller/descriptor.bin

	rm -rf gen
	$(DOCKER_RUN_CMD) prototool all api
	$(DOCKER_RUN_CMD) gofmt -w -s .
	$(DOCKER_RUN_CMD) goimports -local github.com/percona-platform/dbaas-api -w .

gen-dev: docker-build                      ## Same as `gen` but with DEV prototool Docker image
	env DOCKER_RUN_IMAGE=$(DOCKER_DEV_IMAGE) make gen

gen-code:                                  ## Generate code
	go generate ./...
	go install ./...

descriptors:                               ## Update files used for breaking changes detection
	$(DOCKER_RUN_CMD) prototool break descriptor-set api/controller -o api/controller/descriptor.bin

docker-build:                              ## Build prototool Docker dev image
	docker build --pull --squash --tag $(DOCKER_DEV_IMAGE) -f Dockerfile .

docker-push:                               ## Tag and push prototool Docker image
	docker tag $(DOCKER_DEV_IMAGE) $(DOCKER_RUN_IMAGE)
	docker push $(DOCKER_RUN_IMAGE)

run-dev:                                   ## Run bash in prototool Docker dev image
	# the same as DOCKER_RUN_CMD but with `-it` and dev image
	docker run -it --rm --mount='type=bind,src=$(PWD),dst=/work' $(DOCKER_DEV_IMAGE) /bin/bash

.PHONY: gen
