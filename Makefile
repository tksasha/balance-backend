APP=cmd/main.go

PACKAGES=\
	github.com/tksasha/balance/config \
	github.com/tksasha/balance/internal/app \
	github.com/tksasha/balance/internal/app/db \
	github.com/tksasha/balance/internal/app/router \
	github.com/tksasha/balance/internal/controller \
	github.com/tksasha/balance/internal/interface/sqlite3 \
	github.com/tksasha/balance/internal/model \
	github.com/tksasha/balance/internal/repository \
	github.com/tksasha/balance/internal/usecase/item \

.PHONY: all
all: test

.PHONY: run
run:
	@go run $(APP)

r: run

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
test: vet fix fmt
	@echo "go test"
	@go test $(PACKAGES)
