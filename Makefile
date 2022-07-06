.PHONY: install
install: ## install packages
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix ./...
