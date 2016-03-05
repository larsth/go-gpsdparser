package gpsdparser

import (
	"github.com/juju/errors"
	"github.com/larsth/go-gpsdfilter"
)

func ruleParse(p *ParseArgs, rule *gpsdfilter.Rule) (interface{}, error) {
	var (
		i   interface{}
		err error
	)
	switch rule.Class {
	case "ATT":
		i, err = parseATT(p)
		return i, errors.Trace(err)
	case "DEVICE":
		i, err = parseDEVICE(p)
		return i, errors.Trace(err)
	case "DEVICES":
		i, err = parseDEVICES(p)
		return i, errors.Trace(err)
	case "ERROR":
		i, err = parseERROR(p)
		return i, errors.Trace(err)
	case "GST":
		i, err = parseGST(p)
		return i, errors.Trace(err)
	case "POLL":
		i, err = parsePOLL(p)
		return i, errors.Trace(err)
	case "PPS":
		i, err = parsePPS(p)
		return i, errors.Trace(err)
	case "SKY":
		i, err = parseATT(p)
		return i, errors.Trace(err)
	case "TOFF":
		parseTOFF(p)
		return i, errors.Trace(err)
	case "TPV":
		i, err = parseTPV(p)
		return i, errors.Trace(err)
	case "VERSION":
		i, err = parseVERSION(p)
		return i, errors.Trace(err)
	case "WATCH":
		i, err = parseWATCH(p)
		return i, errors.Trace(err)
	}
	return nil, errors.Annotatef(ErrUnknownClass,
		"%s: \"class\"=\"<something>\"",
		`In the gpsd JSON document this JSON value exists: `)
}
