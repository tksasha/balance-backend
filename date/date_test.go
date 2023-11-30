package date_test

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/tksasha/balance/assert"
	"github.com/tksasha/balance/date"
)

var exp = date.Date(time.Date(2023, 11, 23, 0, 0, 0, 0, time.UTC))

func TestNew(t *testing.T) {
	res := date.New(2023, 11, 23)

	Eq(t, res, exp)
}

func TestParse(t *testing.T) {
	t.Run("when `input` is valid", func(t *testing.T) {
		res, err := date.Parse("2023-11-23")

		Eq(t, err, nil)
		Eq(t, res, exp)
	})

	t.Run("when `input` is not valid", func(t *testing.T) {
		_, err := date.Parse("")

		Ne(t, err, nil)
	})
}

func TestToday(t *testing.T) {
	year, month, day := time.Now().Date()

	exp := date.New(year, int(month), day)

	res := date.Today()

	Eq(t, res, exp)
}

func TestIsZero(t *testing.T) {
	t.Run("when it is not empty", func(t *testing.T) {
		subject := date.New(2023, 11, 23)

		Eq(t, subject.IsZero(), false)
	})

	t.Run("when it is empty", func(t *testing.T) {
		subject, _ := date.Parse("")

		Eq(t, subject.IsZero(), true)
	})
}

func TestString(t *testing.T) {
	res := date.New(2023, 11, 23).String()

	Eq(t, res, "2023-11-23")
}

func TestMarshalJSON(t *testing.T) {
	subject, err := date.New(2023, 11, 23).MarshalJSON()

	Eq(t, err, nil)
	Eq(t, string(subject), `"2023-11-23"`)

	t.Run("when it is in a struct", func(t *testing.T) {
		exp := `{"date":"2023-11-23"}`

		str := struct {
			Date date.Date `json:"date"`
		}{date.New(2023, 11, 23)}

		res, err := json.Marshal(str)

		Eq(t, err, nil)
		Eq(t, string(res), exp)
	})
}

func TestUnmarshalJSON(t *testing.T) {
	res := new(date.Date)

	err := res.UnmarshalJSON([]byte(`"2023-11-23"`))

	Eq(t, err, nil)
	Eq(t, *res, exp)

	t.Run("when it is in a struct", func(t *testing.T) {
		exp := struct {
			Date date.Date `json:"date"`
		}{date.New(2023, 11, 23)}

		res := struct {
			Date date.Date `json:"date"`
		}{}

		err := json.Unmarshal([]byte(`{"date":"2023-11-23"}`), &res)

		Eq(t, err, nil)
		Eq(t, res, exp)
	})
}
