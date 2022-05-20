VERSION = `git rev-parse --short HEAD`
BUILDTIME = `date +%FT%T`
LDFLAGS = "-s -w -X crgo/cmd.Version=$(VERSION) -X crgo/cmd.BuildTime=$(BUILDTIME)"

grpcserve:
  go run -ldflags $(LDFLAGS) main.go grpcserve

httpserve:
  go run -ldflags $(LDFLAGS) main.go httpserve

%:
	@true


codegen:
	sh scripts/codegen.sh