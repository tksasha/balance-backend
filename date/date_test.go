package date_test

import (
	"encoding/json"
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"github.com/tksasha/balance/date"
)

var exp = date.Date(time.Date(2023, 11, 23, 0, 0, 0, 0, time.UTC))

func TestNew(t *testing.T) {
	res := date.New(2023, 11, 23)

	assert.Equal(t, res, exp)
}

func TestParse(t *testing.T) {
	t.Run("when `input` is valid", func(t *testing.T) {
		res, err := date.Parse("2023-11-23")

		assert.NilError(t, err)

		assert.Equal(t, res, exp)
	})

	t.Run("when `input` is not valid", func(t *testing.T) {
		_, err := date.Parse("")

		assert.Assert(t, err != nil)
	})
}

func TestToday(t *testing.T) {
	year, month, day := time.Now().Date()

	exp := date.New(year, int(month), day)

	res := date.Today()

	assert.Equal(t, res, exp)
}

func TestIsZero(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		dat := date.New(2023, 11, 23)

		assert.Assert(t, dat.IsZero() == false)
	})

	t.Run("when it is empty", func(t *testing.T) {
		dat, _ := date.Parse("")

		assert.Assert(t, dat.IsZero())
	})
}

func TestString(t *testing.T) {
	res := date.New(2023, 11, 23).String()

	assert.Equal(t, res, "2023-11-23")
}

func TestMarshalJSON(t *testing.T) {
	dat := date.New(2023, 11, 23)

	res, err := dat.MarshalJSON()

	assert.NilError(t, err)

	assert.Equal(t, string(res), `"2023-11-23"`)

	t.Run("when it is in a struct", func(t *testing.T) {
		exp := `{"date":"2023-11-23"}`

		str := struct {
			Date date.Date `json:"date"`
		}{date.New(2023, 11, 23)}

		res, err := json.Marshal(str)

		assert.NilError(t, err)

		assert.Equal(t, string(res), exp)
	})
}

func TestUnmarshalJSON(t *testing.T) {
	res := new(date.Date)

	err := res.UnmarshalJSON([]byte(`"2023-11-23"`))

	assert.NilError(t, err)

	assert.Equal(t, *res, exp)

	t.Run("when it is in a struct", func(t *testing.T) {
		exp := struct {
			Date date.Date `json:"date"`
		}{date.New(2023, 11, 23)}

		res := struct {
			Date date.Date `json:"date"`
		}{}

		err := json.Unmarshal([]byte(`{"date":"2023-11-23"}`), &res)

		assert.NilError(t, err)

		assert.Equal(t, res, exp)
	})
}
