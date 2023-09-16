package surge

import (
	"github.com/zapsaang/conf_factory/utils/consts"
)

type Parser func(ws *WorkSpace, currentSection string, lines []string) error

var SpecificParser = map[string]Parser{
	consts.SurgeParserKV:     KVParser,
	consts.SurgeParserList:   ListParser,
	consts.SurgeParserRule:   RuleParser,
	consts.SurgeParserHost:   HostParser,
	consts.SurgeParserModule: ModuleParser,
}

var SectionDefaultParser = map[string]Parser{
	consts.SurgeSectionGeneral:       KVParser,
	consts.SurgeSectionReplica:       KVParser,
	consts.SurgeSectionSSIDSetting:   KVParser,
	consts.SurgeSectionProxy:         ListParser,
	consts.SurgeSectionProxyGroup:    ListParser,
	consts.SurgeSectionRule:          RuleParser,
	consts.SurgeSectionHost:          HostParser,
	consts.SurgeSectionURLRewrite:    ListParser,
	consts.SurgeSectionHeaderRewrite: ListParser,
	consts.SurgeSectionMITM:          KVParser,
	consts.SurgeSectionScript:        ListParser,
	consts.SurgeSectionModule:        ModuleParser,
}

func ListParser(ws *WorkSpace, currentSection string, lines []string) error {
	currentSectionConent, _ := ws.surgeConfig.Load(currentSection)
	for _, line := range lines {
		currentSectionConent.Store(line, struct{}{})
	}
	return nil
}

func KVParser(ws *WorkSpace, currentSection string, lines []string) error {
	return nil
}

func HostParser(ws *WorkSpace, currentSection string, lines []string) error {
	return nil
}

func RuleParser(ws *WorkSpace, currentSection string, lines []string) error {
	return nil
}

func ModuleParser(ws *WorkSpace, currentSection string, lines []string) error {
	return nil
}
