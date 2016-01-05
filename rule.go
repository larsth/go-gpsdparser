package gpsdparser

import "github.com/larsth/go-gpsdfilter"

func ruleLogIgnoreUnknown(p *ParseArgs, rule *gpsdfilter.Rule) {
	if rule.DoLog {
		p.Logger.Printf("gpsd JSON document of type:\"%s\": %v", rule.Class, p.Data)
	}
	if rule.Type == gpsdfilter.TypeIgnore || rule.Type == gpsdfilter.TypeUnknown {
		return
	}
	if rule.Type == gpsdfilter.TypeLog && rule.DoLog == false {
		p.Logger.Printf("gpsd JSON document of type:\"%s\": %v", rule.Class, p.Data)
	}
	return
}

func ruleParse(p *ParseArgs, rule *gpsdfilter.Rule) (interface{}, error) {
	var (
		i   interface{}
		err error
	)
	switch rule.Class {
	case "ATT":
		parseATT(p)
		return i, err
	case "DEVICE":
		parseDEVICE(p)
		return i, err
	case "DEVICES":
		parseDEVICES(p)
		return i, err
	case "ERROR":
		parseERROR(p)
		return i, err
	case "GST":
		parseGST(p)
		return i, err
	case "POLL":
		parsePOLL(p)
		return i, err
	case "PPS":
		parsePPS(p)
		return i, err
	case "SKY":
		parseATT(p)
		return i, err
	case "TOFF":
		parseTOFF(p)
		return i, err
	case "TPV":
		parseTPV(p)
		return i, err
	case "VERSION":
		parseVERSION(p)
		return i, err
	case "WATCH":
		parseWATCH(p)
		return i, err
	}
	return nil, ErrUnknownClass
}
