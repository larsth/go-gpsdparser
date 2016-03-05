package gpsdparser

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func parseATT(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(*gpsdjson.ATT)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "ATT parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err, "ATT parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}
