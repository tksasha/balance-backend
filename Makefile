.PHONY: all
all: test

.PHONY: test
test:
	GOENV=test go test

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

.PHONY: db
db:
	sqlite3 db/development.sqlite3 '.schema --indent --nosys' > db/schema.sql
	rm -rf db/test.sqlite3
	sqlite3 db/test.sqlite3 '.read db/schema.sql'
