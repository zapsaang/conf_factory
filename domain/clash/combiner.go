package clash

import (
	"reflect"

	"github.com/zapsaang/conf_factory/utils/consts"
)

type Combiner func(p interface{}, ap interface{}) interface{}

var ProfileCombiner map[string]Combiner = map[string]Combiner{
	consts.ClashProfileItemPort:                 nil,
	consts.ClashProfileItemSocksPort:            nil,
	consts.ClashProfileItemMixedPort:            nil,
	consts.ClashProfileItemAllowLan:             nil,
	consts.ClashProfileItemMode:                 nil,
	consts.ClashProfileItemLogLevel:             nil,
	consts.ClashProfileItemExternalController:   nil,
	consts.ClashProfileItemExperimental:         nil,
	consts.ClashProfileItemDns:                  nil,
	consts.ClashProfileItemCFWLatencyTimeout:    nil,
	consts.ClashProfileItemCFWLatencyUrl:        nil,
	consts.ClashProfileItemCFWConnBreakStrategy: nil,
	consts.ClashProfileItemScript:               nil,
	consts.ClashProfileItemRuleProviders:        nil,
	consts.ClashProfileItemCFWBypass:            StringListCombiner,
	consts.ClashProfileItemProxies:              nil,
	consts.ClashProfileItemProxyGroups:          nil,
	consts.ClashProfileItemRules:                StringListCombiner,
}

func SwitchCombiner(p interface{}, ap interface{}) interface{} {
	x, xOK := p.(bool)
	y, yOK := ap.(bool)
	if xOK && yOK {
		return x || y
	}
	if xOK {
		return x
	}
	if yOK {
		return y
	}
	return false
}

func StringCombiner(p interface{}, ap interface{}) interface{} {
	x, xOK := p.(string)
	y, yOK := ap.(string)
	if xOK && yOK {
		if len(y) == 0 {
			return x
		}
		return y
	}
	if xOK {
		return x
	}
	if yOK {
		return y
	}
	return ""
}

func NumberCombiner(p interface{}, ap interface{}) interface{} {
	x, xOK := p.(int)
	y, yOK := ap.(int)
	if xOK && yOK {
		if y == 0 {
			return x
		}
		return y
	}
	if xOK {
		return x
	}
	if yOK {
		return y
	}
	return 0
}

func StringListCombiner(p interface{}, ap interface{}) interface{} {
	pv := reflect.ValueOf(p)
	apv := reflect.ValueOf(ap)
	if pv.Kind() == reflect.Slice && apv.Kind() == reflect.Slice &&
		((pv.Len() > 0 && pv.Index(0).Kind() == reflect.String) || (apv.Len() > 0 && apv.Index(0).Kind() == reflect.String)) {
		pvLen := pv.Len()
		apvLen := apv.Len()
		result := make([]string, pvLen+apvLen)
		for i := 0; i < len(result); i++ {
			if i >= pvLen {
				result[i] = apv.Index(i - pvLen).String()
				continue
			}
			result[i] = pv.Index(i).String()
		}
		return result
	}

	if pv.Kind() == reflect.Slice {
		return p
	}
	if apv.Kind() == reflect.Slice {
		return ap
	}
	return []interface{}{}
}
