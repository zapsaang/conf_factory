package surge

import (
	"bytes"
	"strings"

	"github.com/zapsaang/conf_factory/pkg/maps"
	"github.com/zapsaang/conf_factory/utils/consts"
)

type Storer func(section *maps.OrderedMap[string, any]) []byte

var SectionDefaultStorer = map[string]Storer{
	consts.SurgeSectionGeneral:       KVStorer,
	consts.SurgeSectionReplica:       KVStorer,
	consts.SurgeSectionSSIDSetting:   KVStorer,
	consts.SurgeSectionProxy:         ListStorer,
	consts.SurgeSectionProxyGroup:    ListStorer,
	consts.SurgeSectionRule:          ListStorer,
	consts.SurgeSectionHost:          ListStorer,
	consts.SurgeSectionURLRewrite:    ListStorer,
	consts.SurgeSectionHeaderRewrite: ListStorer,
	consts.SurgeSectionMITM:          KVStorer,
	consts.SurgeSectionScript:        ListStorer,
}

func KVStorer(section *maps.OrderedMap[string, any]) []byte {
	var sectionBuf bytes.Buffer
	section.Range(func(key string, value any) bool {
		sectionBuf.WriteString(key)
		sectionBuf.WriteRune(consts.SurgeKVSeparator)
		_value, ok := value.(maps.OrderedMap[string, struct{}])
		if !ok {
			return true
		}
		var v = make([]string, 0, _value.Len())
		_value.Range(func(key string, value struct{}) bool {
			v = append(v, key)
			return true
		})
		sectionBuf.WriteString(strings.Join(v, consts.SurgeKVJoiner))
		sectionBuf.WriteRune('\n')
		return true
	})
	return sectionBuf.Bytes()
}

func ListStorer(section *maps.OrderedMap[string, any]) []byte {
	var sectionBuf bytes.Buffer
	section.Range(func(key string, value any) bool {
		sectionBuf.WriteString(key)
		sectionBuf.WriteRune('\n')
		return true
	})
	return sectionBuf.Bytes()
}
