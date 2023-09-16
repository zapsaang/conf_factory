package logs

import (
	"context"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/zapsaang/conf_factory/utils/consts"
)

var logger Logger

func init() {
	logger = Logger{
		logrus.New(),
	}
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.DebugLevel)
}

func SetRotateLogs(dir string) {
	if len(dir) == 0 {
		logger.Info("empty rotate log dir")
		return
	}
	rotateLogsName := dir + consts.RotateLogsName
	latestLogLinkName := dir + consts.LatestLogLinkName
	writer, err := rotatelogs.New(rotateLogsName,
		rotatelogs.WithLinkName(latestLogLinkName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logger.WithError(err).
			WithField("rotate_log_dir", dir).
			WithField("rotate_logs_name", rotateLogsName).
			WithField("latest_log_link_name", latestLogLinkName).
			Error("new rotate log writer failed")
		return
	}
	logger.SetOutput(writer)
}

func SetLevel(logLevel string) {
	logger.SetLevel(parseLevel(logLevel))
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *logrus.Entry {
	return logger.WithField(logrus.ErrorKey, err)
}

// WithContext creates an entry from the standard logger and adds a context to it.
func WithContext(ctx context.Context) *logrus.Entry {
	return logger.WithContext(ctx)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

// WithTime creates an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *logrus.Entry {
	return logger.WithTime(t)
}

// Tracef logs a message at level Trace on the standard logger.
func Trace(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debug(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Info(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warn(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Error(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panic(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
