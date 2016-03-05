package gpsdparser

import (
	"github.com/juju/errors"
)

func checkParseArgs(p *ParseArgs) error {
	if p == nil {
		return errors.Trace(ErrNilParseArgs)
	}
	if p.Data == nil {
		return errors.Trace(ErrNilByteSlice)
	}
	if p.Filter == nil {
		return errors.Trace(ErrNilFilter)
	}
	return nil
}
