package gpsdparser

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func parseERROR(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(*gpsdjson.ERROR)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "ERROR parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"ERROR parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}
