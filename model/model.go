package model

import (
	"github.com/go-playground/validator/v10"
)

type Model struct {
	Errors Errors `json:"-"`
}

type Validatable interface {
	AddError(string, string)
}

func (model *Model) IsValid() bool {
	return model.Errors.IsEmpty()
}

func (model *Model) IsNotValid() bool {
	return model.Errors.IsNotEmpty()
}

func (model *Model) AddError(attribute, message string) {
	model.Errors.Add(attribute, message)
}

func Validate(validatable Validatable) {
	validate := validator.New()

	if errs := validate.Struct(validatable); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			attribute := err.Field()

			message := err.Tag()

			validatable.AddError(attribute, message)
		}
	}
}