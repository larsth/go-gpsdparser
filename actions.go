package gpsdparser

import (
	"encoding/json"

	"github.com/larsth/go-gpsdjson"
)

func parseATT(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(*gpsdjson.ATT)
		err error
	)
	
	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}

	return (interface{})(v), nil
}

func parseDEVICE(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(*gpsdjson.DEVICE)
		err error
	)
	
	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}

	return (interface{})(v), nil
}

func parseDEVICES(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(*gpsdjson.DEVICES)
		err error
	)
	
	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}

	return (interface{})(v), nil
}

func parseERROR(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(*gpsdjson.ERROR)
		err error
	)
	
	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}

	return (interface{})(v), nil
}

func parseGST(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.GST)
		err error
	)

	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}
	
	return (interface{})(v), nil
}

func parsePOLL(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.POLL)
		err error
	)

	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}
	
	return (interface{})(v), nil
}

func parsePPS(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.PPS)
		err error
	)

	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}
	
	return (interface{})(v), nil
}

func parseSKY(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.SKY)
		err error
	)

	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}
	
	return (interface{})(v), nil
}

func parseTOFF(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.TOFF)
		err error
	)

	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}
	
	return (interface{})(v), nil
}

func parseTPV(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.TPV)
		err error
	)

	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}
	
	return (interface{})(v), nil
}

func parseVERSION(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.VERSION)
		err error
	)

	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}
	
	return (interface{})(v), nil
}

func parseWATCH(p *ParseArgs) (interface{}, error) {
	var (
		v   = new(gpsdjson.WATCH)
		err error
	)

	if p == nil {
		p.Logger.Println(ErrNilByteSlice.Error())
		return nil, ErrNilByteSlice
	}
	if err = json.Unmarshal(p.Data, v); err != nil {
		p.Logger.Println(err.Error())
		return nil, err
	}
	
	return (interface{})(v), nil
}
