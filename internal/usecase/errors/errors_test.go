package errors_test

import (
	"testing"

	"github.com/tksasha/balance/internal/usecase/errors"
	"gotest.tools/v3/assert"
)

func TestRecordInvalidError(t *testing.T) {
	assert.Error(t, errors.RecordInvalidError, "Record Invalid Error")
}

func TestUnknownError(t *testing.T) {
	assert.Error(t, errors.UnknownError, "Unknown Error")
}
