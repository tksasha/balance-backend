APP=cmd/balance/main.go
PACKAGES=\
	github.com/tksasha/balance/config \
	github.com/tksasha/balance/internal/app \
	github.com/tksasha/balance/internal/app/router \
	github.com/tksasha/balance/internal/interfaces/api \
	github.com/tksasha/balance/internal/interfaces/sqlite3 \
	github.com/tksasha/balance/internal/interfaces/sqlite3/category \
	github.com/tksasha/balance/internal/interfaces/sqlite3/item \
	github.com/tksasha/balance/internal/interfaces/testdb/category \
	github.com/tksasha/balance/internal/interfaces/testdb/errors \
	github.com/tksasha/balance/internal/interfaces/testdb/item \
	github.com/tksasha/balance/internal/models \
	github.com/tksasha/balance/internal/usecases/category \

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
