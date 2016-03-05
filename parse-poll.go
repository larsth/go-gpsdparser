package gpsdparser

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func parsePOLL(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.POLL)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "POLL parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"POLL parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}