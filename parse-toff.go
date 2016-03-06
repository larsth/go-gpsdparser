package gpsdparser

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func parseTOFF(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.TOFF)
		err error
	)

	if p == nil {
		annotatedErr := errors.Trace(ErrNilParseArgs)
		return nil, annotatedErr
	}

	if len(p.Data) == 0 {
		annotatedErr := errors.Trace(ErrEmptyOrNilByteSlice)
		return nil, annotatedErr
	}

	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"TOFF parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}
