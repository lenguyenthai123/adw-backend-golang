package constant

import "errors"

var (
	ErrEncodingToken        = errors.New("error encoding token")
	ErrInvalidToken         = errors.New("invalid token")
	ErrCannotMarshalPayload = errors.New("cannot marshal payload")
	ErrCannotUnmarshalToken = errors.New("cannot unmarshal token")
)
