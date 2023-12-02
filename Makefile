PACKAGE=github.com/tksasha/balance
PACKAGES=$(PACKAGE) \
				 github.com/tksasha/balance/date \
				 github.com/tksasha/balance/formula \
				 github.com/tksasha/balance/model \
				 github.com/tksasha/balance/utils/strings \

.PHONY: all
all: t

.PHONY: t
t:
	@GOENV=test go test $(PACKAGES)

.PHONY: r
r:
	go run $(PACKAGE)

.PHONY: tr
tr:
	GOENV=test make r

.PHONY: db
db:
	@sqlite3 db/development.sqlite3 '.schema --indent --nosys' > db/schema.sql
	@rm -rf db/test.sqlite3
	@sqlite3 db/test.sqlite3 '.read db/schema.sql'
