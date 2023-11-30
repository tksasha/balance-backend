package strings_test

import (
	"testing"

	. "github.com/tksasha/balance/assert"
	"github.com/tksasha/balance/utils/strings"
)

func TestToSnakeCase(t *testing.T) {
	t.Run("when it is `HelloWorld`", func(t *testing.T) {
		subject := strings.ToSnakeCase("HelloWorld")

		Eq(t, subject, "hello_world")
	})

	t.Run("when it is `CategoryID`", func(t *testing.T) {
		subject := strings.ToSnakeCase("CategoryID")

		Eq(t, subject, "category_id")
	})

	t.Run("when it is `СлаваУкраїні`", func(t *testing.T) {
		subject := strings.ToSnakeCase("СлаваУкраїні")

		Eq(t, subject, "слава_україні")
	})
}
