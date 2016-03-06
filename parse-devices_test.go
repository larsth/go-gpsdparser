package gpsdparser

import (
	"testing"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdjson"
)

func TestParseDevicesNilParseArgs(t *testing.T) {
	var wantErr = errors.Trace(ErrNilParseArgs)
	utilsParseWithNilParseArgs(wantErr, t, parseDEVICES)
}

func TestParseDevicesEmptyOrNilByteSlice(t *testing.T) {
	var wantErr = errors.Trace(ErrEmptyOrNilByteSlice)
	utilsParseWithEmptyOrNilByteSlice(wantErr, t, parseDEVICES)
}

func TestParseDevicesJsonUnmarshalError(t *testing.T) {
	const annotation = "DEVICES parsing"
	var (
		p       = []byte(`{"class":"DEVICES"`)
		wantErr = errors.Annotate(
			wantParseJsonUnmarshalErr, annotation)
	)
	utilsParseWithJsonUnmarshalError(p, wantErr, t, parseDEVICES)
}

func TestParseDevices(t *testing.T) {
	var (
		pA              *ParseArgs = new(ParseArgs)
		p                          = []byte(`{"class":"DEVICES"}`)
		gotI            interface{}
		gotGpsdJsonType *gpsdjson.DEVICES
		gotErr          error
		wantErr         error = nil
		s               string
		ok              bool
	)

	pA.Data = p

	gotI, gotErr = parseDEVICES(pA)

	//error test
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}

	//type conversion test
	gotGpsdJsonType, ok = (gotI).(*gpsdjson.DEVICES)
	if ok {
		t.Logf("\nType conversion ok\n\tgot:\n\t\t%#v\n", gotGpsdJsonType)
	} else {
		sGpsdjsonType := `*gpsdjson.DEVICES`
		t.Logf(
			"gotI' of type 'interface{}' is not type convertable to type '%s'",
			sGpsdjsonType)
		t.Errorf("gotI' of type 'interface{}' is '%#v'", gotGpsdJsonType)
	}
}
