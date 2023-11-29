package model

import (
	"github.com/tksasha/balance/utils/strings"
)

type Errors map[string]map[string][]string

func NewErrors(args ...string) Errors {
	errs := Errors{"errors": {}}

	if len(args) == 2 {
		errs.Add(args[0], args[1])
	}

	return errs
}

func (errs Errors) Add(attribute, message string) {
	attribute = strings.ToSnakeCase(attribute)

	errs["errors"][attribute] = append(errs["errors"][attribute], message)
}

func (errs Errors) IsEmpty() bool {
	return errs.size() == 0
}

func (errs Errors) size() int {
	return len(errs["errors"])
}
