package item

type repository struct{}

func New() *repository {
	return &repository{}
}
