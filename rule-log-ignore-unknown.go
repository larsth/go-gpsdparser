package gpsdparser

import (
	"fmt"

	"github.com/larsth/go-gpsdfilter"
)

func ruleLogIgnoreUnknown(p *ParseArgs, rule *gpsdfilter.Rule) string {
	if rule.DoLog {
		s := fmt.Sprintf("gpsd JSON document of type:\"%s\": %v",
			rule.Class, p.Data)
		return s
	}
	if rule.Type == gpsdfilter.TypeIgnore || rule.Type == gpsdfilter.TypeUnknown {
		return ""
	}
	if rule.Type == gpsdfilter.TypeLog && rule.DoLog == false {
		s := fmt.Sprintf("gpsd JSON document of type:\"%s\": %v",
			rule.Class, p.Data)
		return s
	}
	return "?"
}
