package consts

import (
	"fmt"
	"regexp"
	"time"
)

const (
	ProgramName         = "conf_enhancement"
	WorkSpacePath       = "/work_space"
	JSONIndent          = "    "
	EmptyString         = ""
	ZeroInt64     int64 = 0

	CurrentDir             = "."
	CurrentParentDir       = ".."
	UnixAbsolutePathPrefix = '/'
	UnixPathSeparator      = '/'

	DefaultSurgeBaseDir   = "base/surge"
	DefaultSurgeResultDir = "result/surge"

	DefaultClashBaseDir   = "base/clash"
	DefaultClashResultDir = "result/clash"

	HTTPRequestTimeout = 10 * time.Second
	FileBufferLength   = 4096
	AntsPoolLimitation = 8

	ParamNameToken = "token"

	TemplateServerAddress = "%s:%d"

	RegexpPatternURL = `^https?://[^\s/$.?#].[^\s]*$`
)

var (
	RotateLogsName    = fmt.Sprintf("/%s.%%Y%%m%%d.log", ProgramName)
	LatestLogLinkName = fmt.Sprintf("/%s.latest.log", ProgramName)
	BytesTabCharacter = []byte{'\t'}
	BytesIndent4Space = []byte{' ', ' ', ' ', ' '}

	RegexpCompiledURL = regexp.MustCompile(RegexpPatternURL)
)
