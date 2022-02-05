export GO111MODULE=on
export GOFLAGS=-mod=vendor

.PHONY: build
build: vendor FORCE
	go build
	go build github.com/bopmatic/sdk/golang/util

.PHONY: test
test: build FORCE
	go test

go.sum vendor: go.mod
	go mod vendor

.PHONY: clean
clean:
	rm -rf go.sum vendor

FORCE:
