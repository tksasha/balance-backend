APP=cmd/app/main.go
PACKAGES=\
	github.com/tksasha/balance/config \
	github.com/tksasha/balance/internal/app \
	github.com/tksasha/balance/internal/app/db \
	github.com/tksasha/balance/internal/app/router \
	github.com/tksasha/balance/internal/controller \
	github.com/tksasha/balance/internal/interface/sqlite3/category \
	github.com/tksasha/balance/internal/interface/sqlite3/item \
	github.com/tksasha/balance/internal/interface/testdb/category \
	github.com/tksasha/balance/internal/interface/testdb/errors \
	github.com/tksasha/balance/internal/interface/testdb/item \
	github.com/tksasha/balance/internal/model \
	github.com/tksasha/balance/internal/repository \
	github.com/tksasha/balance/internal/usecase \
	github.com/tksasha/balance/internal/usecase/errors \
	github.com/tksasha/balance/pkg/model \
	github.com/tksasha/balance/pkg/model/errors \

.PHONY: all
all: vet fix fmt test

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
