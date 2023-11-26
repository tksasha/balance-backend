.PHONY: all
all: tests

.PHONY: tests
tests:
	@GOENV=test go test

.PHONY: t
t: tests

.PHONY: integrations
integrations: build db
	@GOENV=test ./balance &
	@sleep 1
	@cd tests && bundle && rspec spec/items/create_spec.rb spec/items/show_spec.rb
	@killall balance
	@go clean

.PHONY: i
i: integrations

.PHONY: run
run:
	@go run \
		$(filter-out $(wildcard *_test.go), $(wildcard *.go))

.PHONY: r
r: run

.PHONY: build
build:
	@go build \
		-o balance \
		$(filter-out $(wildcard *_test.go), $(wildcard *.go))

.PHONY: b
b: build

.PHONY: db
db:
	@sqlite3 db/development.sqlite3 '.schema --indent --nosys' > db/schema.sql
	@rm -rf db/test.sqlite3
	@sqlite3 db/test.sqlite3 '.read db/schema.sql'
