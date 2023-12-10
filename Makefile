GOVET=find . -name *.go -exec go vet {} \;
GOFIX=find . -name *.go -exec go fix {} \;
GOFMT=find . -name *.go -exec gofmt -l -s -w {} \;
GOTEST=find . -name *_test.go -exec go test {} \;
GOLINT=golangci-lint run
APP=cmd/app/main.go

.PHONY: all
all: vet fix fmt lint test

.PHONY: run
run:
	@go run $(APP)

.PHONY: vet
vet:
	@echo "go vet"
	@$(GOVET)

.PHONY: fix
fix:
	@echo "go fix"
	@$(GOFIX)

.PHONY: fmt
fmt:
	@echo "go fmt"
	@$(GOFMT)

.PHONY: test
test:
	@echo "go test"
	@$(GOTEST) $(PACKAGES)

.PHONY: lint
lint:
	@echo "go lint"
	@$(GOLINT)
