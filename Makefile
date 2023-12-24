GOFMT=find . -name *.go -exec gofmt -l -s -w {} \;
GOLINT=golangci-lint run
APP=cmd/app/main.go
PACKAGES= \
	github.com/tksasha/balance/cmd/app \
	github.com/tksasha/balance/config \
	github.com/tksasha/balance/internal/app \
	github.com/tksasha/balance/internal/interfaces/api \
	github.com/tksasha/balance/internal/interfaces/sqlite3 \
	github.com/tksasha/balance/internal/interfaces/sqlite3/dummy \
	github.com/tksasha/balance/internal/models \
	github.com/tksasha/balance/internal/repositories \
	github.com/tksasha/balance/internal/usecases \

.PHONY: all
all: vet fix fmt lint test

.PHONY: run
run:
	@go run $(APP)

.PHONY: vet
vet:
	@echo "go vet"
	@go vet $(PACKAGES)

.PHONY: fix
fix:
	@echo "go fix"
	@go fix $(PACKAGES)

.PHONY: fmt
fmt:
	@echo "go fmt"
	@$(GOFMT)

.PHONY: test
test:
	@echo "go test"
	@go test $(PACKAGES)

.PHONY: lint
lint:
	@echo "go lint"
	@$(GOLINT)
