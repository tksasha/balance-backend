.PHONY: run all

all:
	@GOENV=test go test

run:
	@go run \
		check_content_type.go \
		consts.go \
		create_item_handler.go \
		db.go \
		item.go \
		item_params.go \
		main.go \
