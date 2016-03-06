package gpsdparser

import (
	"testing"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func TestParsePollNilParseArgs(t *testing.T) {
	var wantErr = errors.Trace(ErrNilParseArgs)
	utilsParseWithNilParseArgs(wantErr, t, parsePOLL)
}

func TestParsePollEmptyOrNilByteSlice(t *testing.T) {
	var wantErr = errors.Trace(ErrEmptyOrNilByteSlice)
	utilsParseWithEmptyOrNilByteSlice(wantErr, t, parsePOLL)
}

func TestParsPollJsonUnmarshalError(t *testing.T) {
	const annotation = "POLL parsing"
	var (
		p       = []byte(`{"class":"POLL"`)
		wantErr = errors.Annotate(
			wantParseJsonUnmarshalErr, annotation)
	)
	utilsParseWithJsonUnmarshalError(p, wantErr, t, parsePOLL)
}

func TestParsePOLL(t *testing.T) {
	var (
		pA              *ParseArgs = new(ParseArgs)
		p                          = []byte(`{"class":"POLL"}`)
		gotI            interface{}
		gotGpsdJsonType *gpsdjson.POLL
		gotErr          error
		wantErr         error = nil
		s               string
		ok              bool
	)

	pA.Data = p

	gotI, gotErr = parsePOLL(pA)

	//error test
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}

	//type conversion test
	gotGpsdJsonType, ok = (gotI).(*gpsdjson.POLL)
	if ok {
		t.Logf("\nType conversion ok\n\tgot:\n\t\t%#v\n", gotGpsdJsonType)
	} else {
		sGpsdjsonType := `*gpsdjson.POLL`
		t.Logf(
			"gotI' of type 'interface{}' is not type convertable to type '%s'",
			sGpsdjsonType)
		t.Errorf("gotI' of type 'interface{}' is '%#v'", gotGpsdJsonType)
	}
}
