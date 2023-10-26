package growatt

import "errors"

var (
	ErrResponseMustNotBeHTML = errors.New("response must not be HTML")
	ErrTooManyRequests       = errors.New("too many request")
)
