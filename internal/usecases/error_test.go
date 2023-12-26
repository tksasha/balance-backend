package usecases_test

import (
	"encoding/json"
	"testing"

	"github.com/tksasha/balance/internal/usecases"
	"gotest.tools/v3/assert"
)

type dbError struct{}

func (err *dbError) Error() string {
	return `{"database":"error"}`
}

func TestError(t *testing.T) {
	expected := `{"application":{"database":"error"}}`

	result, err := json.Marshal(
		usecases.NewError(
			&dbError{},
		),
	)

	assert.NilError(t, err)
	assert.Equal(t, string(result), expected)
}
