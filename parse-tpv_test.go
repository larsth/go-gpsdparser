package gpsdparser

import (
	"testing"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func TestParseTPVNilParseArgs(t *testing.T) {
	var wantErr = errors.Trace(ErrNilParseArgs)
	utilsParseWithNilParseArgs(wantErr, t, parseTPV)
}

func TestParseTPVEmptyOrNilByteSlice(t *testing.T) {
	var wantErr = errors.Trace(ErrEmptyOrNilByteSlice)
	utilsParseWithEmptyOrNilByteSlice(wantErr, t, parseTPV)
}

func TestParsTPVJsonUnmarshalError(t *testing.T) {
	const annotation = "TPV parsing"
	var (
		p       = []byte(`{"class":"TPV"`)
		wantErr = errors.Annotate(
			wantParseJsonUnmarshalErr, annotation)
	)
	utilsParseWithJsonUnmarshalError(p, wantErr, t, parseTPV)
}

func TestParseTPV(t *testing.T) {
	var (
		pA              *ParseArgs = new(ParseArgs)
		p                          = []byte(`{"class":"TPV"}`)
		gotI            interface{}
		gotGpsdJsonType *gpsdjson.TPV
		gotErr          error
		wantErr         error = nil
		s               string
		ok              bool
	)

	pA.Data = p

	gotI, gotErr = parseTPV(pA)

	//error test
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}

	//type conversion test
	gotGpsdJsonType, ok = (gotI).(*gpsdjson.TPV)
	if ok {
		t.Logf("\nType conversion ok\n\tgot:\n\t\t%#v\n", gotGpsdJsonType)
	} else {
		sGpsdjsonType := `*gpsdjson.TPV`
		t.Logf(
			"gotI' of type 'interface{}' is not type convertable to type '%s'",
			sGpsdjsonType)
		t.Errorf("gotI' of type 'interface{}' is '%#v'", gotGpsdJsonType)
	}
}
