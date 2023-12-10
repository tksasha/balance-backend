LINT=golangci-lint run

APP=cmd/app/app.go
PACKAGES=\

.PHONY: all
all: vet fix fmt lint test

.PHONY: run
run:
	@go run $(APP)

.PHONY: vet
vet:
	@echo "go vet"
	@go vet $(APP)
	@go vet $(PACKAGES)

.PHONY: fix
fix:
	@echo "go fix"
	@go fix $(APP)
	@go fix $(PACKAGES)

.PHONY: fmt
fmt:
	@echo "go fmt"
	@go fmt $(APP)
	@go fmt $(PACKAGES)

.PHONY: test
test:
	@echo "go test"
	@go test $(PACKAGES)

.PHONY: lint
lint:
	@echo "go lint"
	@$(LINT)
