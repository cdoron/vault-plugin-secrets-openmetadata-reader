GOARCH = amd64
OS = linux

.DEFAULT_GOAL := all

all: build

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS="$(OS)" GOARCH="$(GOARCH)" go build -o vault/plugins/vault-plugin-secrets-openmetadata-reader cmd/vault-plugin-secrets-openmetadata-reader/main.go

.PHONY: enable
enable:
	vault secrets enable -path=openmetadata-secrets-reader vault-plugin-secrets-openmetadata-reader 

.PHONY: clean
clean:
	rm -f ./vault/plugins/vault-plugin-secrets-openmetadata-reader

.PHONY: test
test:
	go test -v ./...

include hack/make-rules/verify.mk
