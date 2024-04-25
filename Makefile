PROJECTNAME=`go list`
VERSION := $(shell git describe --tag --abbrev=0)
NOW := $(shell date -u +'%Y%m%d-%H%M%S')


build:
	@go build -o bin/terraform-provider-http2b64

lint: fmt
	@golangci-lint run

fmt:
	@gofmt -s -w -e .

test:
	@go test -v -cover -timeout=120s -parallel=4 ./...

testacc:
	TF_ACC=1 go test -v -cover -timeout 120m ./...

generate:
	@go generate

commit: fmt generate test
	@git add .
	@git commit -a -m "Update $(NOW)"
	@git push

server:
	@python3 -m http.server --directory testdata 80