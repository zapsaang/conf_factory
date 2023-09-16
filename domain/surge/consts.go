package surge

import (
	"regexp"
	"strings"

	"github.com/zapsaang/conf_factory/utils/consts"
)

const (
	ReTitleWithSpecificParser = `\[\s*([^@\]]+)\s*@\s*([^@\]]+)\s*\]`
	ReSpecificParser          = `\s*@\s*([^@\]]+)\s*`
)

var (
	RegexpTitleWithSpecificParser = regexp.MustCompile(ReTitleWithSpecificParser)
	RegexpSpecificParser          = regexp.MustCompile(ReSpecificParser)

	SectionGeneralLowerTitle       = strings.ToLower(consts.SurgeSectionGeneral)
	SectionReplicaLowerTitle       = strings.ToLower(consts.SurgeSectionReplica)
	SectionSSIDSettingLowerTitle   = strings.ToLower(consts.SurgeSectionSSIDSetting)
	SectionProxyLowerTitle         = strings.ToLower(consts.SurgeSectionProxy)
	SectionProxyGroupLowerTitle    = strings.ToLower(consts.SurgeSectionProxyGroup)
	SectionRuleLowerTitle          = strings.ToLower(consts.SurgeSectionRule)
	SectionHostLowerTitle          = strings.ToLower(consts.SurgeSectionHost)
	SectionURLRewriteLowerTitle    = strings.ToLower(consts.SurgeSectionURLRewrite)
	SectionHeaderRewriteLowerTitle = strings.ToLower(consts.SurgeSectionHeaderRewrite)
	SectionMITMLowerTitle          = strings.ToLower(consts.SurgeSectionMITM)
	SectionScriptLowerTitle        = strings.ToLower(consts.SurgeSectionScript)
	SectionModuleLowerTitle        = strings.ToLower(consts.SurgeSectionModule)

	OrderedSurgeTitle = []string{
		consts.SurgeSectionGeneral,
		consts.SurgeSectionReplica,
		consts.SurgeSectionSSIDSetting,
		consts.SurgeSectionProxy,
		consts.SurgeSectionProxyGroup,
		consts.SurgeSectionRule,
		consts.SurgeSectionHost,
		consts.SurgeSectionURLRewrite,
		consts.SurgeSectionHeaderRewrite,
		consts.SurgeSectionMITM,
		consts.SurgeSectionScript,
		consts.SurgeSectionModule,
	}

	ProcessFlow = []ProcessNode{
		Load,
		Integrate,
		Filte,
		Store,
	}
)
