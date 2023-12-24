package sqlite3_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tksasha/balance/internal/interfaces/sqlite3"
	"gotest.tools/v3/assert"
)

func TestNewQueryError(t *testing.T) {
	expected := `{"database":{"query":"unknown column"}}`

	result, err := json.Marshal(
		sqlite3.NewQueryError(
			fmt.Errorf("unknown column"), //nolint:goerr113
		),
	)

	assert.NilError(t, err)
	assert.Equal(t, string(result), expected)
}
