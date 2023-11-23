package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("GOENV") != "test" {
		panic("What are You doing?!")
	}

	code := func(m *testing.M) int {
		db := Open()

		truncate(db)

		defer func() {
			truncate(db)

			Close(db)
		}()

		return m.Run()
	}(m)

	os.Exit(code)
}

func truncate(db *sql.DB) {
	for _, table := range []string{"categories", "cashes", "items"} {
		db.Exec(fmt.Sprintf("DELETE FROM %s", table))
	}
}
