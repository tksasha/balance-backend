package assert

import (
	"testing"
)

func In[S ~[]E, E comparable] (t *testing.T, s S, e E) {
	for _, c := range s {
		if c == e {
			return
		}
	}

	t.Errorf("\033[31m`%v` doesn't include `%v`\033[0m", s, e)
}
