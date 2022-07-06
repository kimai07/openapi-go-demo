.PHONY: install
install: ## install packages
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix ./...

.PHONY: codegen
codegen: ## generate code
	rm -f api/generated/openapi/*.gen.go
	oapi-codegen -generate "chi-server" -old-config-style -package openapi openapi/api/openapi.yaml > api/generated/openapi/chi-server.gen.go
	oapi-codegen -generate "spec" -old-config-style -package openapi openapi/api/openapi.yaml > api/generated/openapi/spec.gen.go
	oapi-codegen -generate "types" -old-config-style -package openapi openapi/api/openapi.yaml > api/generated/openapi/types.gen.go
