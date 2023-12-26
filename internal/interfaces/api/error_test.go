package api_test

import (
	"encoding/json"
	"testing"

	"github.com/tksasha/balance/internal/interfaces/api"
	"gotest.tools/v3/assert"
)

type appError struct{}

func (err appError) Error() string {
	return `{"application":"error"}`
}

func TestError(t *testing.T) {
	expected := `{"api":{"application":"error"}}`

	result, err := json.Marshal(
		api.NewError(&appError{}),
	)

	assert.NilError(t, err)
	assert.Equal(t, string(result), expected)
}
