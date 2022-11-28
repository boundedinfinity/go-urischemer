package urischemer

import "github.com/boundedinfinity/go-commoner/errorer"

var (
	ErrUriSchemeNotFound  = errorer.Errorf("schema not found")
	ErrUriSchemeNotFoundv = errorer.Errorfn(ErrUriSchemeNotFound)
)
