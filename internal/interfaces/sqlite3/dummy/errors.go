package dummy

import (
	"fmt"
)

var (
	ErrNotFound = fmt.Errorf("NOT FOUND")
	ErrUnknown  = fmt.Errorf("UNKNOWN")
)
