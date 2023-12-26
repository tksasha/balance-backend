package dummy

type NotFoundError struct{}

func (err *NotFoundError) Error() string {
	return `{"database":"not found"}`
}

type UnknownError struct{}

func (err *UnknownError) Error() string {
	return `{"database":"unknown"}`
}
