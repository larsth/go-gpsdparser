package gpsdparser

import (
	"github.com/juju/errors"
)

func checkParseArgs(p *ParseArgs) error {
	if p == nil {
		return errors.Trace(ErrNilParseArgs)
	}
	if len(p.Data) == 0 {
		return errors.Trace(ErrEmptyOrNilByteSlice)
	}
	if p.Filter == nil {
		return errors.Trace(ErrNilFilter)
	}
	return nil
}
