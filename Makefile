BINARY:=tink-reboot
BUILD_ARGS:=GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags '-s -w -extldflags "-static"'

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## complie for linux
	GOOS=linux ${BUILD_ARGS} -o bin/${BINARY}-linux-amd64 main.go

.PHONY: image
image: ## build the container image
	docker build -t tink-rebooter:local . 

.PHONY: lint
lint:  ## run linting
	@echo be sure golangci-lint is installed: https://golangci-lint.run/usage/install/
	golangci-lint run

.PHONY: goimports
goimports: ## run goimports
	@echo be sure goimports is installed
	goimports -w ./

.PHONY: goimports-check
goimports-check: ## run goimports displaying diffs
	@echo be sure goimports is installed
	goimports -d . | (! grep .)
