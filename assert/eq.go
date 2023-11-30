package assert

import (
	"testing"
)

func Eq[E comparable] (t *testing.T, l, r E) {
	if l != r {
		t.Errorf("\033[31m`%v` != `%v`\033[0m", l, r)
	}
}
