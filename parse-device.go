package gpsdparser

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func parseDEVICE(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.DEVICE)
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
			"DEVICE parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}
