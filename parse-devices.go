package gpsdparser

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func parseDEVICES(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(*gpsdjson.DEVICES)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "DEVICES parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"DEVICES parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}
