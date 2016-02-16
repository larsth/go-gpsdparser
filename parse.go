//Package gpsdparser parses gpsd JSON documents into types from package
//"github.com/larsth/go-gpsdjson".
package gpsdparser

import "github.com/larsth/go-gpsdfilter"
import "github.com/juju/errors"

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
}

func checkParseArgs(p *ParseArgs) error {
	if p == nil {
		return errors.Trace(ErrNilParseArgs)
	}
	if p.Data == nil {
		return errors.Trace(ErrNilByteSlice)
	}
	if p.Filter == nil {
		return errors.Trace(ErrNilFilter)
	}
	return nil
}

//Parse parses a gpsd JSON document with the arguments from the ParseArgs
func Parse(p *ParseArgs) (interface{}, error) {
	var (
		rule              *gpsdfilter.Rule
		err, annotatedErr error
		i                 interface{}
	)

	if err = checkParseArgs(p); err != nil {
		annotatedErr = errors.Annotate(err, "checkParseArgs error")
		return nil, annotatedErr
	}

	if rule, err = p.Filter.Filter(p.Data); err != nil {
		annotatedErr = errors.Annotate(err, "filter error")
		return nil, annotatedErr
	}

	ruleLogIgnoreUnknown(p, rule)
	if rule.Type == gpsdfilter.TypeParse {
		if i, err = ruleParse(p, rule); err != nil {
			annotatedErr = errors.Annotate(err, "ruleParse error")
			return nil, annotatedErr
		}
	}
	return i, nil
}
