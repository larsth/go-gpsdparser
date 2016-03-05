package gpsdparser

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func parseWATCH(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.WATCH)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "WATCH parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"WATCH parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}