package usecases_test

import (
	"encoding/json"
	"testing"

	"github.com/tksasha/balance/internal/usecases"
	"gotest.tools/v3/assert"
)

type DBError struct{}

func (err *DBError) Error() string {
	return `{"database":"error"}`
}

func TestApplicationError(t *testing.T) {
	expected := `{"application":{"database":"error"}}`

	result, err := json.Marshal(
		usecases.NewApplicationError(
			&DBError{},
		),
	)

	assert.NilError(t, err)
	assert.Equal(t, string(result), expected)
}
