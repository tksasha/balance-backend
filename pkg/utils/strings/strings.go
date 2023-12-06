package strings

import (
	"regexp"
	"strings"
)

// Converts `HelloWorld` to `hello_world`
// or `CategoryID` to `category_id`
// even with UTF support
// `СлаваУкраїні` to `слава_україні`
//
// if you don't want to support UTF
// use `regexp.MustCompile("([a-z])([A-Z])")` instead
func ToSnakeCase(input string) string {
	input = regexp.
		MustCompile("(\\P{Lu}+)(\\p{Lu}+)").
		ReplaceAllString(input, "${1}_${2}")

	return strings.ToLower(input)
}
