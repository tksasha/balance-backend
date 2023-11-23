package model

type Model struct {
	Errors Errors `json:"-"`
}

func (model *Model) IsValid() bool {
	return model.Errors.IsEmpty()
}

func (model *Model) IsNotValid() bool {
	return model.Errors.IsNotEmpty()
}
