GOPATH:=$(shell go env GOPATH)

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: run
run:
	@go run main.go

.PHONY: build
build:
	@go build -o *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker-build
docker-build:
	@DOCKER_BUILDKIT=1 docker build -t simple-gin-rest:latest .

.PHONY: docker-run
docker-run:
	@docker run -p 8080:8080 simple-gin-rest:latest