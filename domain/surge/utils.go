package surge

import (
	"strings"

	"github.com/zapsaang/conf_factory/utils/consts"
)

func IsSectionTitle(line string) bool {
	lineLen := len(line)
	return lineLen > 2 && line[0] == '[' && line[lineLen-1] == ']'
}

func GetSectionParser(title string) (string, Parser) {
	if !RegexpTitleWithSpecificParser.MatchString(title) {
		return getSectionDefaultParser(title)
	}

	parserName := strings.ToLower(strings.TrimSpace(RegexpSpecificParser.FindString(title)))
	pureTitle := GetSectionStdTitle(RegexpSpecificParser.ReplaceAllString(title, consts.ConstEmptyString))
	parser, ok := SpecificParser[parserName]
	if ok {
		return pureTitle, parser
	}
	return getSectionDefaultParser(pureTitle)
}

func getSectionDefaultParser(title string) (string, Parser) {
	parser, ok := SectionDefaultParser[title]
	if ok {
		return title, parser
	}
	return title, ListParser
}

func GetSectionStorer(title string) (Storer, bool) {
	storer, ok := SectionDefaultStorer[title]
	if ok {
		return storer, true
	}
	return nil, false
}

func GetSectionStdTitle(title string) string {
	switch strings.ToLower(title) {
	case SectionGeneralLowerTitle:
		return consts.SurgeSectionGeneral
	case SectionReplicaLowerTitle:
		return consts.SurgeSectionReplica
	case SectionSSIDSettingLowerTitle:
		return consts.SurgeSectionSSIDSetting
	case SectionProxyLowerTitle:
		return consts.SurgeSectionProxy
	case SectionProxyGroupLowerTitle:
		return consts.SurgeSectionProxyGroup
	case SectionRuleLowerTitle:
		return consts.SurgeSectionRule
	case SectionHostLowerTitle:
		return consts.SurgeSectionHost
	case SectionURLRewriteLowerTitle:
		return consts.SurgeSectionURLRewrite
	case SectionHeaderRewriteLowerTitle:
		return consts.SurgeSectionHeaderRewrite
	case SectionMITMLowerTitle:
		return consts.SurgeSectionMITM
	case SectionScriptLowerTitle:
		return consts.SurgeSectionScript
	case SectionModuleLowerTitle:
		return consts.SurgeSectionModule
	default:
		return title
	}
}
