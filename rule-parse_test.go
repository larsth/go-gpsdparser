package gpsdparser

import (
	"testing"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdfilter"
)

func TestRuleParseNilParseArgs(t *testing.T) {
	var (
		pA      *ParseArgs       // is 'nil'
		r       *gpsdfilter.Rule // is 'nil'
		wantErr = errors.Trace(ErrNilParseArgs)
		gotErr  error
		s       string
		ok      bool
	)
	_, gotErr = ruleParse(pA, r)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestRuleParseNilRule(t *testing.T) {
	var (
		pA      = &ParseArgs{}
		r       *gpsdfilter.Rule // is 'nil'
		wantErr = errors.Trace(ErrNilRule)
		gotErr  error
		s       string
		ok      bool
	)

	_, gotErr = ruleParse(pA, r)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestRuleParseUnknown(t *testing.T) {
	var (
		pA      = &ParseArgs{}
		p       = []byte(`{"class":"krakelikrox"}`)
		r       *gpsdfilter.Rule
		wantErr error

		gotErr error
		s      string
		ok     bool
	)

	pA.Data = p
	r = &gpsdfilter.Rule{
		Class: `krakelikrox`,
		Type:  gpsdfilter.TypeParse,
		DoLog: true}
	wantErr = errors.Annotatef(ErrUnknownClass,
		"gpsd document JSON type '%s' is unknown to this Go package",
		r.Class)

	_, gotErr = ruleParse(pA, r)

	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

type testTableNoErrorsRuleParse struct {
	P     []byte
	Class string
}

var testDataNoErrorsRuleParse = []*testTableNoErrorsRuleParse{
	//Test 0:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"ATT"}`),
		Class: "ATT",
	},
	//Test 1:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"DEVICE"}`),
		Class: "DEVICE",
	},
	//Test 2:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"DEVICES"}`),
		Class: "DEVICES",
	},
	//Test 3:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"ERROR"}`),
		Class: "ERROR",
	},
	//Test 4:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"GST"}`),
		Class: "GST",
	},
	//Test 5:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"POLL"}`),
		Class: "POLL",
	},
	//Test 6:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"PPS"}`),
		Class: "PPS",
	},
	//Test 7:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"TOFF"}`),
		Class: "TOFF",
	},
	//Test 8:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"TPV"}`),
		Class: "TPV",
	},
	//Test 9:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"VERSION"}`),
		Class: "VERSION",
	},
	//Test 10:
	&testTableNoErrorsRuleParse{
		P:     []byte(`{"class":"WATCH"}`),
		Class: "WATCH",
	},
}

func TestRuleParseNoErrors(t *testing.T) {
	var (
		wantErr error // is 'nil' - always!
		pA      *ParseArgs
		r       *gpsdfilter.Rule // is 'nil'
		gotErr  error
		s       string
		ok      bool
	)

	for i, td := range testDataNoErrorsRuleParse {
		pA = new(ParseArgs)
		pA.Data = td.P
		//pA.Filter = &gpsdfilter.Filter{} //gpsdfilter.New()

		r = &gpsdfilter.Rule{Class: td.Class}
		r.Type = gpsdfilter.TypeParse
		r.DoLog = false

		_, gotErr = ruleParse(pA, r)

		s, ok = errorTestI(gotErr, wantErr, i, "testDataNoErrorsRuleParse")
		if false == ok {
			t.Error(s)
		}
	}
}
