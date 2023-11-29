package jsonapi

import (
	"encoding/json"
)

var details = map[string]string{
	"server error": "Internal Server Error",
	"invalid":      "Is not valid",
	"not found":    "Resource not found",
	"required":     "Can't be blank",
}

type source struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

type err struct {
	Title  string  `json:"title"`
	Detail string  `json:"detail"`
	Source *source `json:"source,omitempty"`
}

type Errors []err

func NewErrors(args ...string) (errs Errors) {
	errs = make(Errors, 0, 0)

	errs.add(args)

	return
}

func (errs *Errors) Add(args ...string) {
	errs.add(args)
}

func (errs *Errors) add(args []string) {
	if len(args) == 1 {
		*errs = append(*errs, err{Title: args[0], Detail: details[args[0]]})

		return
	}

	if len(args) < 3 {
		return
	}

	src, attr, title := args[0], args[1], args[2]

	er := err{Title: title}

	if len(args) > 3 {
		er.Detail = args[3]
	} else {
		er.Detail = details[title]
	}

	switch src {
	case "pointer":
		er.Source = &source{Pointer: attr}
	case "parameter":
		er.Source = &source{Parameter: attr}
	default:
		return
	}

	*errs = append(*errs, er)
}

func (errs Errors) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string][]err{"errors": errs})
}

func (errs Errors) IsEmpty() bool {
	return len(errs) == 0
}
