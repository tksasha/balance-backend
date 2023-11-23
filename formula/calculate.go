package formula

import (
	"go/token"
	"go/types"
	"regexp"
	"strconv"
)

func Calculate(formula string) (result float64, err error) {
	formula =
		regexp.
			MustCompile("[^0-9+-/*()]").
			ReplaceAllString(formula, "")

	fset := token.NewFileSet()

	var pos token.Pos

	pkg := types.NewPackage("", "")

	res, err := types.Eval(fset, pkg, pos, formula)

	if err != nil {
		return 0, err
	}

	result, err = strconv.ParseFloat(res.Value.String(), 64)

	return
}
