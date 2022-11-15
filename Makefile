PB_FILE_LIST_FILE     := config/pb_file_list.txt
PB_FILE_LIST :=$(shell cat $(PB_FILE_LIST_FILE))
ifndef   PB_FILE_LIST
	PB_FILE_LIST=*
endif


GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
API_PROTO_FILES=$(shell find api -name "$(PB_FILE_LIST).proto")
KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)

COMPILE_TARGET="./"


$(info )
$(info KRATOS version:$(KRATOS))
$(info make file list:$(PB_FILE_LIST))
$(info )

.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go get -u github.com/google/wire/cmd/wire@v0.5.0
	go get -u github.com/envoyproxy/protoc-gen-validate@v0.6.1

.PHONY: grpc
# generate grpc code
grpc:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:$(COMPILE_TARGET) \
		--go-grpc_out=paths=source_relative:$(COMPILE_TARGET) \
		$(API_PROTO_FILES)

.PHONY: http
# generate http code
http:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:$(COMPILE_TARGET)  \
		--go-http_out=paths=source_relative:$(COMPILE_TARGET)  \
		$(API_PROTO_FILES)

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:$(COMPILE_TARGET)  \
           --hierrors_out=paths=source_relative:$(COMPILE_TARGET)  \
           $(API_PROTO_FILES)

.PHONY: validate
# generate validate code
validate:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:$(COMPILE_TARGET)  \
           --validate_out=paths=source_relative,lang=go:$(COMPILE_TARGET) \
           $(API_PROTO_FILES)

.PHONY: proto
# generate internal proto
proto:
	protoc --proto_path=. \
		--proto_path=./third_party \
 		--go_out=paths=source_relative:$(COMPILE_TARGET)  \
		$(API_PROTO_FILES)

.PHONY: swagger
# generate swagger file
swagger:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--openapiv2_out ./swagger \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt json_names_for_fields=false \
		$(API_PROTO_FILES)

.PHONY: generate
# generate client code
generate:
	cd $(COMPILE_TARGET) && go generate ./...

.PHONY: build
# build mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...
build:

.PHONY: test
# test
test:

.PHONY: all
# generate all
all:
	#make generate;
	make grpc;
	make http;
	make errors;
	make validate;
	make proto;
	make swagger;
	#make build;
	#make test;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
