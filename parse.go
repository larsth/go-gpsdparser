//Package gpsdparser parses gpsd JSON documents into types from package
//"github.com/larsth/go-gpsdjson".
package gpsdparser

import (
	"log"

	"github.com/larsth/go-gpsdfilter"
)

//ParseArgs is arguments for function Parse.
//
//Data is a JSON document.
//
//Filter is filter, which tells the parser what
//to do with gpsd JSON document.
//
//Logger is logger used to log errors, and log gpsd JSON documents.
type ParseArgs struct {
	Data   []byte
	Filter *gpsdfilter.Filter
	Logger *log.Logger
}

func checkParseArgs(p *ParseArgs) error {
	if p == nil {
		return ErrNilParseArgs
	}
	if p.Data == nil {
		return ErrNilByteSlice
	}
	if p.Filter == nil {
		return ErrNilFilter
	}
	if p.Logger == nil {
		return ErrNilLogger
	}
	return nil
}

//Parse parses a gpsd JSON document with the arguments from the ParseArgs
func Parse(p *ParseArgs) (interface{}, error) {
	var (
		rule *gpsdfilter.Rule
		err  error
		i    interface{}
	)

	if err = checkParseArgs(p); err != nil {
		return nil, err
	}

	if rule, err = p.Filter.Filter(p.Data); err != nil {
		return nil, err
	}

	ruleLogIgnoreUnknown(p, rule)
	if rule.Type == gpsdfilter.TypeParse {
		if i, err = ruleParse(p, rule); err != nil {
			return nil, err
		}
	}
	return i, err
}
