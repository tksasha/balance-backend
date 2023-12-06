package category

type repository struct{}

func New() *repository {
	return &repository{}
}
