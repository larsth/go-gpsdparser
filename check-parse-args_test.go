package gpsdparser

import (
	"testing"

	"github.com/juju/errors"
	"github.com/larsth/go-gpsdfilter"
)

func TestCheckParseArgsNilArgs(t *testing.T) {
	var (
		pA      *ParseArgs // is 'nil'
		wantErr = errors.Trace(ErrNilParseArgs)
		gotErr  error
		s       string
		ok      bool
	)
	gotErr = checkParseArgs(pA)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestCheckParseArgsEmptyOrNilByteSlice(t *testing.T) {
	var (
		pA      = new(ParseArgs)
		p       []byte // is 'nil' == len(p) is 0, zero
		wantErr = errors.Trace(ErrEmptyOrNilByteSlice)
		gotErr  error
		s       string
		ok      bool
	)
	pA.Data = p
	gotErr = checkParseArgs(pA)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestCheckParseArgsNilFilter(t *testing.T) {
	var (
		pA      = new(ParseArgs)
		p       = []byte(`{"class":"ATT"}`)
		f       *gpsdfilter.Filter // is 'nil'
		wantErr = errors.Trace(ErrNilFilter)
		gotErr  error
		s       string
		ok      bool
	)
	pA.Data = p
	pA.Filter = f
	gotErr = checkParseArgs(pA)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestCheckParseArgsNoErrors(t *testing.T) {
	var (
		pA      = new(ParseArgs)
		p       = []byte(`{"class":"ATT"}`)
		f       = &gpsdfilter.Filter{}
		wantErr error // is 'nil'
		gotErr  error
		s       string
		ok      bool
	)
	pA.Data = p
	pA.Filter = f
	gotErr = checkParseArgs(pA)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}
