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

func parseDEVICE(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(*gpsdjson.DEVICE)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "DEVICE parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"DEVICE parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}

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

func parseGST(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.GST)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "GST parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"GST parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}

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

func parsePPS(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.PPS)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "PPS parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"PPS parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}

func parseSKY(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.SKY)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "SKY parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"SKY parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}

func parseTOFF(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.TOFF)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "TOFF parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"TOFF parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}

func parseTPV(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.TPV)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "TPV parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"TPV parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}

func parseVERSION(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.VERSION)
		err error
	)

	if p == nil {
		annotatedErr := errors.Annotate(ErrNilByteSlice, "VERSION parse error")
		return nil, annotatedErr
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		annotatedErr := errors.Annotate(err,
			"VERSION parsing: json.Unmarshal error")
		return nil, annotatedErr
	}

	return (interface{})(v), nil
}

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
