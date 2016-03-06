package gpsdparser

import (
	"testing"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdfilter"
)

func TestParseCheckParseArgsError(t *testing.T) {
	var (
		pA      *ParseArgs // is 'nil'
		wantErr = errors.New(
			`checkParseArgs error: The ParseArgs is a nil pointer`)
		gotErr error
		s      string
		ok     bool
	)
	_, gotErr = Parse(pA)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestParseFilterError(t *testing.T) {
	var (
		pA *ParseArgs
		//r       *gpsdfilter.Rule
		wantErr = errors.New(
			`filter error: filter rule not found: Filter: No such Rule`)
		gotErr error
		s      string
		ok     bool
	)

	//	r = &gpsdfilter.Rule{
	//		Class: `TPV`,
	//		DoLog: false,
	//		Type:  gpsdfilter.TypeParse}

	pA = new(ParseArgs)
	pA.Data = []byte(`{"class":"TPV"}`)
	pA.Filter = gpsdfilter.New()
	//pA.Filter.Add(r)

	_, gotErr = Parse(pA)

	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestParseRuleParseError(t *testing.T) {
	var (
		pA      *ParseArgs
		r       *gpsdfilter.Rule
		wantErr = errors.New(
			`ruleParse error: gpsd document JSON type 'krakelikrox' ` +
				`is unknown to this Go package: Unknown gpsd JSON document ` +
				`class type`)
		gotErr error
		s      string
		ok     bool
	)

	r = &gpsdfilter.Rule{
		Class: `krakelikrox`,
		DoLog: false,
		Type:  gpsdfilter.TypeParse}

	pA = new(ParseArgs)
	pA.Data = []byte(`{"class":"krakelikrox"}`)
	pA.Filter = gpsdfilter.New()
	pA.Filter.Add(r)

	_, gotErr = Parse(pA)

	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestParseRuleNoErrors(t *testing.T) {
	var (
		pA           *ParseArgs
		r            *gpsdfilter.Rule
		wantErr      error // is 'nil'
		gotInterface interface{}
		gotErr       error
		s            string
		ok           bool
	)

	r = &gpsdfilter.Rule{
		Class: `TPV`,
		DoLog: false,
		Type:  gpsdfilter.TypeParse}

	pA = new(ParseArgs)
	pA.Data = []byte(`{"class":"TPV"}`)
	pA.Filter = gpsdfilter.New()
	pA.Filter.Add(r)

	gotInterface, gotErr = Parse(pA)

	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
	if gotInterface == nil {
		t.Error(`'interface{}' returned from Parse() is nil, ` +
			`want interface{} != nil`)
	}
}
