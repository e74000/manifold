package manifold

import "errors"

var (
	ErrorGETFailed             = errors.New("GET failed")
	ErrorPOSTFailed            = errors.New("POST failed")
	ErrorFailedToParseResponse = errors.New("failed to parse response")
)
