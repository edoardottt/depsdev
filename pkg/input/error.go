package input

import "errors"

var (
	ErrArgumentLeast         = errors.New("argument must be specified")
	ErrArgumentsLeast        = errors.New("arguments must be specified")
	ErrInvalidPackageManager = errors.New("invalid package manager specified")
)
