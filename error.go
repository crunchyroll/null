package null

import "errors"

var (
	ErrUnsupportedValue = errors.New(`This value does not support unmarshaling.`)
)
