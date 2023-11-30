package formula_test

import (
	"testing"

	. "github.com/tksasha/balance/assert"
	"github.com/tksasha/balance/formula"
)

func TestCalculateFormula(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		res, err := formula.Calculate("2+2")

		Eq(t, err, nil)
		Eq(t, res, 4.0)
	})

	t.Run("when it contains spaces", func(t *testing.T) {
		res, err := formula.Calculate(" 2 + 2 ")

		Eq(t, err, nil)
		Eq(t, res, 4.0)
	})

	t.Run("when it contains letters", func(t *testing.T) {
		res, err := formula.Calculate("a2b+c2d")

		Eq(t, err, nil)
		Eq(t, res, 4.0)
	})

	t.Run("when it is an addition", func(t *testing.T) {
		res, err := formula.Calculate("3+4")

		Eq(t, err, nil)
		Eq(t, res, 7.0)
	})

	t.Run("when it is a substraction", func(t *testing.T) {
		res, err := formula.Calculate("8-6")

		Eq(t, err, nil)
		Eq(t, res, 2.0)
	})

	t.Run("when it is a multiplication", func(t *testing.T) {
		res, err := formula.Calculate("4*5")

		Eq(t, err, nil)
		Eq(t, res, 20.0)
	})

	t.Run("when it is a division", func(t *testing.T) {
		res, err := formula.Calculate("9/3")

		Eq(t, err, nil)
		Eq(t, res, 3.0)
	})

	t.Run("when it is empty", func(t *testing.T) {
		res, err := formula.Calculate("")

		Ne(t, err, nil)
		Eq(t, res, 0.0)
	})

	t.Run("when it contains brackets", func(t *testing.T) {
		res, err := formula.Calculate("(5 + 3) * 4")

		Eq(t, err, nil)
		Eq(t, res, 32.0)
	})

	t.Run("when calculated result is float64", func(t *testing.T) {
		res, err := formula.Calculate("42.1 + 69.01")

		Eq(t, err, nil)
		Eq(t, res, 111.11)
	})
}
