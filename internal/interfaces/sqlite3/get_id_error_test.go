package sqlite3_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tksasha/balance/internal/interfaces/sqlite3"
	"gotest.tools/v3/assert"
)

func TestNewGetIDError(t *testing.T) {
	expected := `{"database":{"GetIDError":"get id error"}}`

	result, err := json.Marshal(
		sqlite3.NewGetIDError(
			fmt.Errorf("get id error"), //nolint:goerr113
		),
	)

	assert.NilError(t, err)
	assert.Equal(t, string(result), expected)
}
