VERSION = `git rev-parse --short HEAD`
BUILDTIME = `date +%FT%T`
LDFLAGS = "-s -w -X crgo/cmd.Version=$(VERSION) -X crgo/cmd.BuildTime=$(BUILDTIME)"


include scripts/make-rules/golang.mk

## serve: run http serve  and grpc serve
.PHONY: serve
serve:
	go run -ldflags $(LDFLAGS) main.go serve



%:
	@true

## codegen: proto file generate
.PHONY: codegen
codegen:
	sh scripts/codegen.sh



.PHONY: tidy
tidy:
	@$(GO) mod tidy

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo  "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
