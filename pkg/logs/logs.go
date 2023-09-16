package logs

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func parseLevel(lvl string) logrus.Level {
	switch strings.ToLower(lvl) {
	case "panic", "silence":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace", "verbose":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}
