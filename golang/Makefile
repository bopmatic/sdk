export GO111MODULE=on
export GOFLAGS=-mod=vendor

OPENAPIGENGO=java -jar ./openapi-generator-cli.jar generate -g go -o openapi --git-host github.com --git-user-id bopmatic --git-repo-id sdk/golang/openapi

.PHONY: build
build: goswag.generated vendor openapi FORCE
	go build
	go build github.com/bopmatic/sdk/golang/util

# https://git.wiki.kernel.org/index.php/GitFaq#Can_I_add_empty_directories.3F
test_assets/project/emptySiteAssets/siteAssetsEmpty:
	mkdir test_assets/project/emptySiteAssets/siteAssetsEmpty

.PHONY: test
test: build test_assets/project/emptySiteAssets/siteAssetsEmpty FORCE
	go test -v

unit-tests.xml: test_assets/project/emptySiteAssets/siteAssetsEmpty FORCE
	gotestsum --junitfile unit-tests.xml

vendor: go.mod
	go mod download
	go mod vendor

.PHONY: deps
deps: goswag.generated openapi
	rm -rf go.mod go.sum vendor openapi-generator-cli.jar
	go mod init github.com/bopmatic/sdk/golang
	GOPROXY=direct go mod tidy

# swagger install via:
# dir=$(mktemp -d) 
# git clone https://github.com/go-swagger/go-swagger "$dir" 
# cd "$dir"
# go mod vendor
# go install ./cmd/swagger

# protoc-gen-openapiv2 install via:
# GOFLAGS='' go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# jq install:
# sudo apt install jq

goswag.generated: pb/sr.bopmatic.json
	swagger generate client -f pb/sr.bopmatic.json -c goswag --default-scheme='https' -A goswag
	touch goswag.generated

pb/sr.pb.go pb/sr_grpc.pb.go pb/sr.swagger.json: pb/sr.proto
	protoc -I ./ --go_out ./ --go_opt paths=source_relative --go-grpc_out ./ --go-grpc_opt paths=source_relative --openapiv2_out=./ --openapiv2_opt=simple_operation_ids=true,generate_unbound_methods=true ./pb/sr.proto

pb/sr.bopmatic.json: pb/sr.swagger.json
	jq '{"host":"api.bopmatic.com"} + .' pb/sr.swagger.json | jq '{"schemes":["https"]} + .' > pb/sr.bopmatic.json

openapi-generator-cli.jar:
	curl -L https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.2.0/openapi-generator-cli-6.2.0.jar --output openapi-generator-cli.jar

.PHONY: openapi
openapi: openapi-generator-cli.jar pb/sr.bopmatic.json
	$(OPENAPIGENGO) -i ./pb/sr.bopmatic.json
	rm openapi/go.mod

.PHONY: clean
clean:
	rm -rf pb/sr.pb.go pb/sr_grpc.pb.go ./test_assets/package/stub/.bopmatic pb/sr.swagger.json pb/sr.bopmatic.json openapi goswag models goswag.generated

FORCE:
