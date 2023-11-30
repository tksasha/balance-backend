package assert

import (
	"errors"
	"testing"
)

func Is(t *testing.T, l, r any) {
	switch l.(type) {
	case error:
		if !errors.Is(l.(error), r.(error)) {
			t.Errorf("\033[31m`%v` is not `%v`\033[0m", l, r)
		}
	default:
		t.Errorf("\033[31m`%v` UNKNOWN TYPE `%v`\033[0m")
	}
}
