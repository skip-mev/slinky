package client

import (
	"errors"
)

// Client sentinel errors.
var (
	ErrorNilRequest = errors.New("request cannot be nil")
)
