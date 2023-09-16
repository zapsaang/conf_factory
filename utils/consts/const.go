package consts

import "fmt"

const (
	ConstProgramName         = "conf_enhancement"
	ConstWorkSpacePath       = "/work_space"
	ConstJSONIndent          = "    "
	ConstEmptyString         = ""
	ConstZeroInt64     int64 = 0

	ConstCurrentDir             = "."
	ConstCurrentParentDir       = ".."
	ConstUnixAbsolutePathPrefix = '/'
	ConstUnixPathSeparator      = '/'

	ConstDefaultSurgeBaseDir   = "base/surge"
	ConstDefaultSurgeResultDir = "result/surge"

	ConstDefaultClashBaseDir   = "base/clash"
	ConstDefaultClashResultDir = "result/clash"

	ParamNameToken = "token"

	TemplateServerAddress = "%s:%d"
)

var (
	RotateLogsName    = fmt.Sprintf("/%s.%%Y%%m%%d.log", ConstProgramName)
	LatestLogLinkName = fmt.Sprintf("/%s.latest.log", ConstProgramName)
)
