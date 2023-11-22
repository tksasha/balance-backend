.PHONY: run all

all:
	@GOENV=test go test

run:
	@go run \
		check_content_type.go \
		create_item_handler.go \
		create_item_query.go \
		db.go \
		item.go \
		item_params.go \
		main.go \
