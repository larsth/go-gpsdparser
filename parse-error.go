package gpsdparser

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func parseERROR(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.ERROR)
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
			"ERROR parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}
