package gpsdparser

import (
	"testing"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func TestParseAttNilParseArgs(t *testing.T) {
	var wantErr = errors.Trace(ErrNilParseArgs)
	utilsParseWithNilParseArgs(wantErr, t, parseATT)
}

func TestParseAttEmptyOrNilByteSlice(t *testing.T) {
	var wantErr = errors.Trace(ErrEmptyOrNilByteSlice)
	utilsParseWithEmptyOrNilByteSlice(wantErr, t, parseATT)
}

func TestParseAttJsonUnmarshalError(t *testing.T) {
	const annotation = "ATT parsing"
	var (
		p       = []byte(`{"class":"ATT"`)
		wantErr = errors.Annotate(
			wantParseJsonUnmarshalErr, annotation)
	)
	utilsParseWithJsonUnmarshalError(p, wantErr, t, parseATT)
}

func TestParseAtt(t *testing.T) {
	var (
		pA              *ParseArgs = new(ParseArgs)
		p                          = []byte(`{"class":"ATT"}`)
		gotI            interface{}
		gotGpsdJsonType *gpsdjson.ATT
		gotErr          error
		wantErr         error = nil
		s               string
		ok              bool
	)

	pA.Data = p

	gotI, gotErr = parseATT(pA)

	//error test
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}

	//type conversion test
	gotGpsdJsonType, ok = (gotI).(*gpsdjson.ATT)
	if ok {
		t.Logf("\nType conversion ok\n\tgot:\n\t\t%#v\n", gotGpsdJsonType)
	} else {
		sGpsdjsonType := `*gpsdjson.ATT`
		t.Logf(
			"gotI' of type 'interface{}' is not type convertable to type '%s'",
			sGpsdjsonType)
		t.Errorf("gotI' of type 'interface{}' is '%#v'", gotGpsdJsonType)
	}
}
