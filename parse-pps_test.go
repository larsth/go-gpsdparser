package gpsdparser

import (
	"testing"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func TestParsePPSNilParseArgs(t *testing.T) {
	var wantErr = errors.Trace(ErrNilParseArgs)
	utilsParseWithNilParseArgs(wantErr, t, parsePPS)
}

func TestParsePPSEmptyOrNilByteSlice(t *testing.T) {
	var wantErr = errors.Trace(ErrEmptyOrNilByteSlice)
	utilsParseWithEmptyOrNilByteSlice(wantErr, t, parsePPS)
}

func TestParsPPSJsonUnmarshalError(t *testing.T) {
	const annotation = "PPS parsing"
	var (
		p       = []byte(`{"class":"PPS"`)
		wantErr = errors.Annotate(
			wantParseJsonUnmarshalErr, annotation)
	)
	utilsParseWithJsonUnmarshalError(p, wantErr, t, parsePPS)
}

func TestParsePPS(t *testing.T) {
	var (
		pA              *ParseArgs = new(ParseArgs)
		p                          = []byte(`{"class":"PPS"}`)
		gotI            interface{}
		gotGpsdJsonType *gpsdjson.PPS
		gotErr          error
		wantErr         error = nil
		s               string
		ok              bool
	)

	pA.Data = p

	gotI, gotErr = parsePPS(pA)

	//error test
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}

	//type conversion test
	gotGpsdJsonType, ok = (gotI).(*gpsdjson.PPS)
	if ok {
		t.Logf("\nType conversion ok\n\tgot:\n\t\t%#v\n", gotGpsdJsonType)
	} else {
		sGpsdjsonType := `*gpsdjson.PPS`
		t.Logf(
			"gotI' of type 'interface{}' is not type convertable to type '%s'",
			sGpsdjsonType)
		t.Errorf("gotI' of type 'interface{}' is '%#v'", gotGpsdJsonType)
	}
}
