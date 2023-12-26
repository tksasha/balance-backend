package usecases_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tksasha/balance/internal/usecases"
	"gotest.tools/v3/assert"
)

func TestUnknownError(t *testing.T) {
	expected := `{"application":{"UnknownError":"unknown"}}`

	result, err := json.Marshal(
		usecases.NewUnknownError(
			fmt.Errorf("unknown"), //nolint:goerr113
		),
	)

	assert.NilError(t, err)
	assert.Equal(t, string(result), expected)
}

func TestNotFoundError(t *testing.T) {
	expected := `{"application":{"NotFoundError":"by id"}}`

	result, err := json.Marshal(
		usecases.NewNotFoundError(
			fmt.Errorf("by id"), //nolint:goerr113
		),
	)

	assert.NilError(t, err)
	assert.Equal(t, string(result), expected)
}
