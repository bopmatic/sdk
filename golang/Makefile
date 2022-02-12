export GO111MODULE=on
export GOFLAGS=-mod=vendor

OPENAPIGENGO=java -jar ./openapi-generator-cli.jar generate -g go -o openapi --git-host github.com --git-user-id bopmatic --git-repo-id sdk/golang/openapi

.PHONY: build
build: vendor openapi FORCE
	go build
	go build github.com/bopmatic/sdk/golang/util

.PHONY: test
test: build FORCE
	go test -v

go.sum vendor: go.mod openapi goswag.generated pb/sr.pb.go pb/sr_grpc.pb.go
	go mod vendor
	go mod tidy

goswag.generated: pb/sr.bopmatic.json
	swagger generate client -f pb/sr.bopmatic.json -c goswag --default-scheme='https' -A goswag
	touch goswag.generated

pb/sr.pb.go pb/sr_grpc.pb.go pb/sr.swagger.json: pb/sr.proto
	protoc -I ./ --go_out ./ --go_opt paths=source_relative --go-grpc_out ./ --go-grpc_opt paths=source_relative --openapiv2_out=./ --openapiv2_opt=simple_operation_ids=true,generate_unbound_methods=true ./pb/sr.proto

pb/sr.bopmatic.json: pb/sr.swagger.json
	jq '{"host":"api.bopmatic.com"} + .' pb/sr.swagger.json | jq '{"schemes":["https"]} + .' > pb/sr.bopmatic.json

openapi-generator-cli.jar:
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.4.0/openapi-generator-cli-5.4.0.jar -O openapi-generator-cli.jar

openapi: openapi-generator-cli.jar pb/sr.bopmatic.json
	$(OPENAPIGENGO) -i ./pb/sr.bopmatic.json
	rm openapi/go.mod

.PHONY: clean
clean:
	rm -rf go.sum vendor pb/sr.pb.go pb/sr_grpc.pb.go ./test_assets/package/stub/.bopmatic pb/sr.swagger.json pb/sr.bopmatic.json openapi goswag models goswag.generated

FORCE:
