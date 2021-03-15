VERSION = `git rev-parse --short HEAD`
BUILDTIME = `date +%FT%T`
LDFLAGS = "-s -w -X crgo/cmd.Version=$(VERSION) -X crgo/cmd.BuildTime=$(BUILDTIME)"

serve: 
	go run -ldflags $(LDFLAGS) main.go serve
%:
	@true


codegen:
	sh scripts/codegen.sh