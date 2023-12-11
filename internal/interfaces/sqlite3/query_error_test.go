package sqlite3_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/interfaces/sqlite3"
	"gotest.tools/v3/assert"
)

func TestQueryError(t *testing.T) {
	subject := sqlite3.NewQueryError(errors.New("some error"))

	assert.Error(t, subject, "[DB][QUERY] SOME ERROR")
}
