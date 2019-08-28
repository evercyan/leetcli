.PHONY: help
all: help
help: Makefile
	@echo
	@echo "build: compile leetcli linux and mac version"
	@echo

build:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o ./bin/leetcli-mac helper.go
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ./bin/leetcli-linux helper.go
