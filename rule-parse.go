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

	if p == nil {
		return nil, errors.Trace(ErrNilParseArgs)
	}

	if rule == nil {
		return nil, errors.Trace(ErrNilRule)
	}

	//Below:
	// 'case "SKY"' is excluded, because a "SKY" gpsd JSON document type is
	// always embedded in another gpsd JSON document type, and for that reason,
	//where doesn't exists a 'parseSKY function.

	switch rule.Class {
	case "ATT":
		i, err = parseATT(p)
	case "DEVICE":
		i, err = parseDEVICE(p)
	case "DEVICES":
		i, err = parseDEVICES(p)
	case "ERROR":
		i, err = parseERROR(p)
	case "GST":
		i, err = parseGST(p)
	case "POLL":
		i, err = parsePOLL(p)
	case "PPS":
		i, err = parsePPS(p)
	case "TOFF":
		parseTOFF(p)
	case "TPV":
		i, err = parseTPV(p)
	case "VERSION":
		i, err = parseVERSION(p)
	case "WATCH":
		i, err = parseWATCH(p)
	default:
		err = errors.Annotatef(ErrUnknownClass,
			"gpsd document JSON type '%s' is unknown to this Go package",
			rule.Class)
	}

	if err != nil {
		return nil, errors.Trace(err)
	}
	return i, nil
}
