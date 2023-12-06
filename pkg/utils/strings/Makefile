.PHONY: all
all: test

.PHONY: vet
vet:
	@echo "go vet"
	@go vet

.PHONY: fix
fix:
	@echo "go fix"
	@go fix

.PHONY: fmt
fmt:
	@echo "go fmt"
	@go fmt

.PHONY: test
test: vet fix fmt
	@echo "go test"
	@go test
