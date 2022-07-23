VERSION = `git rev-parse --short HEAD`
BUILDTIME = `date +%FT%T`
LDFLAGS = "-s -w -X crgo/infra.Version=$(VERSION) -X crgo/infra.BuildTime=$(BUILDTIME)"


include scripts/make-rules/golang.mk

## serve: run http serve  and grpc serve
.PHONY: serve
serve:
	go run -ldflags $(LDFLAGS) main.go serve


%:
	@true

## codegen: proto file generate
.PHONY: codegen
codegen: protoc-gen-go
	sh scripts/codegen.sh


.PHONY: protoc-gen-go
protoc-gen-go:
	go get github.com/golang/protobuf/protoc-gen-go@v1.3.5


.PHONY: tidy
tidy:
	@$(GO) mod tidy

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo  "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
