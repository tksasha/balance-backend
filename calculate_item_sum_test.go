package main

import (
	"slices"
	"strings"
	"testing"
)

func TestCalculateItemSum(t *testing.T) {
	t.Run("when `Formula` is empty", func(t *testing.T) {
		item := NewItem()

		item.Formula = ""

		exp := 0.0

		item.Calculate()

		res := item.Sum

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when `Formula` is not valid", func(t *testing.T) {
		item := NewItem()

		item.Formula = "(42.1 + 69.01"

		item.Calculate()

		exp := "is not valid"

		errs := item.Errors.Get("formula")

		res := strings.Join(errs, ", ")

		if !slices.Contains(errs, exp) {
			t.Errorf(M, exp, res)
		}
	})
}
