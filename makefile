cnf ?= config.env
include $(cnf)
build:
	rm -rf ./bin && go build -o bin/terra.bin cmd/terra/main.go

publish:
