.PHONY: all
all: test

.PHONY: tests
tests:
	GOENV=test go test

.PHONY: t
t: tests

.PHONY: integrations
integrations: build
	@GOENV=test ./balance &
	@sleep 1
	@cd tests && bundle && rspec spec/items/create_spec.rb
	@killall balance
	@go clean

.PHONY: i
i: integrations

.PHONY: run
run:
	@go run \
		check_content_type.go \
		create_item_handler.go \
		create_item_query.go \
		db.go \
		item.go \
		item_params.go \
		main.go \

.PHONY: r
r: run

.PHONY: build
build:
	@go build \
		-o balance \
		check_content_type.go \
		create_item_handler.go \
		create_item_query.go \
		db.go \
		item.go \
		item_params.go \
		main.go \

.PHONY: db
db:
	sqlite3 db/development.sqlite3 '.schema --indent --nosys' > db/schema.sql
	rm -rf db/test.sqlite3
	sqlite3 db/test.sqlite3 '.read db/schema.sql'
