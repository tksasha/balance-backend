package model

type Errors map[string]map[string][]string

func NewErrors() Errors {
	return Errors{"errors": {}}
}

func (errs Errors) Add(attribute string, message string) {
	errs["errors"][attribute] = append(errs["errors"][attribute], message)
}

func (errs Errors) Get(attribute string) (messages []string) {
	return errs["errors"][attribute]
}

func (errs Errors) Size() int {
	return len(errs["errors"])
}

func (errs Errors) IsEmpty() bool {
	return errs.Size() == 0
}

func (errs Errors) IsNotEmpty() bool {
	return errs.Size() > 0
}
