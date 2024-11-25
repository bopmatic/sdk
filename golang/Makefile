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
	go test -v github.com/bopmatic/sdk/golang/util

unit-tests.xml: test_assets/project/emptySiteAssets/siteAssetsEmpty FORCE
	gotestsum --junitfile unit-tests.xml

vendor: go.mod
	go mod download
	go mod vendor

.PHONY: deps
deps: grpc-gateway goswag.generated openapi
	rm -rf go.mod go.sum vendor openapi-generator-cli.jar pb/protoc-gen-openapiv2 grpc-gateway
	go mod init github.com/bopmatic/sdk/golang
	GOPROXY=direct go mod tidy

pb/protoc-gen-openapiv2/options/annotations.proto: grpc-gateway
	mkdir -p pb/protoc-gen-openapiv2/options
	cp grpc-gateway/protoc-gen-openapiv2/options/*.proto pb/protoc-gen-openapiv2/options

grpc-gateway:
	git clone https://github.com/grpc-ecosystem/grpc-gateway.git

# swagger install via:
# dir=$(mktemp -d) 
# git clone https://github.com/go-swagger/go-swagger "$dir" 
# cd "$dir"
# go mod vendor
# go install ./cmd/swagger

# protoc-gen-openapiv2 install via:
# GOFLAGS='' go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# protoc install via:
# wget https://github.com/protocolbuffers/protobuf/releases/download/v28.3/protoc-28.3-linux-x86_64.zip
# unzip protoc-28.3-linux-x86_64.zip
# sudo mv bin/protoc /usr/local/bin
# sudo mv include/google /usr/local/include

# protoc-gen-go install via:
# GOFLAGS="" go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# jq install:
# sudo apt install jq

# mockgen install:
# GOFLAGS="" go install github.com/golang/mock/mockgen@v1.6.0

goswag.generated: pb/sr.bopmatic.json pb/signup.bopmatic.json
	swagger generate client -f pb/sr.bopmatic.json -c goswag --default-scheme='https' -A goswag
	swagger generate client -f pb/signup.bopmatic.json -c goswag.signup --default-scheme='https' -A goswag
	touch goswag.generated

pb/sr.pb.go pb/sr_grpc.pb.go pb/sr.swagger.json: pb/sr.proto pb/protoc-gen-openapiv2/options/annotations.proto
	protoc -I ./ -I ./pb -I /usr/local/include --go_out ./ --go_opt paths=source_relative --go-grpc_out ./ --go-grpc_opt paths=source_relative --openapiv2_out=./ --openapiv2_opt=simple_operation_ids=true,generate_unbound_methods=true ./pb/sr.proto

pb/signup.pb.go pb/signup_grpc.pb.go pb/signup.swagger.json: pb/signup.proto pb/protoc-gen-openapiv2/options/annotations.proto
	protoc -I ./ -I ./pb -I /usr/local/include --go_out ./ --go_opt paths=source_relative --go-grpc_out ./ --go-grpc_opt paths=source_relative --openapiv2_out=./ --openapiv2_opt=simple_operation_ids=true,generate_unbound_methods=true ./pb/signup.proto

pb/sr.bopmatic.json: pb/sr.swagger.json
	jq '{"host":"api.bopmatic.com"} + .' pb/sr.swagger.json | jq '{"schemes":["https"]} + .' > pb/sr.swagger.json2
	jq '.security |= [{"BearerAuth": [], "ApiKeyAuth": []}]' pb/sr.swagger.json2 > pb/sr.bopmatic.json

pb/signup.bopmatic.json: pb/signup.swagger.json
	jq '{"host":"api.bopmatic.com"} + .' pb/signup.swagger.json | jq '{"schemes":["https"]} + .' > pb/signup.bopmatic.json

openapi-generator-cli.jar:
	curl -L https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.7.0/openapi-generator-cli-7.7.0.jar --output openapi-generator-cli.jar

.PHONY: openapi
openapi: openapi-generator-cli.jar pb/sr.bopmatic.json
	$(OPENAPIGENGO) -i ./pb/sr.bopmatic.json
	rm openapi/go.mod

.PHONY: clean
clean:
	rm -rf pb/sr.pb.go pb/sr_grpc.pb.go ./test_assets/package/stub/.bopmatic pb/sr.swagger.json pb/sr.swagger.json2 pb/sr.bopmatic.json pb/signup.bopmatic.json openapi goswag goswag.signup models goswag.generated

FORCE:
