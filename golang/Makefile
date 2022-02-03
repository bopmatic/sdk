export GO111MODULE=on
export GOFLAGS=-mod=vendor

.PHONY: build
build: vendor FORCE
	go build

.PHONY: test
test: build FORCE
	go test

go.sum vendor: go.mod
	go mod vendor

.PHONY: clean
clean:
	rm -rf go.sum vendor

FORCE:
