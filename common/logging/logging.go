package logging

import (
  "github.com/Sirupsen/logrus"
  "github.ibm.com/AdvancedAnalyticsCanada/service-common-lib/common/util"
)

var (
	loggerEntry *logrus.Entry
	contextLogger *logrus.Logger
  g_appName string
  DebugFlag bool
)

//@params
//
//@return
//
func init() {
  g_appName = util.GetENV("APP")

	contextLogger = logrus.New()
	contextLogger.Formatter = new(logrus.JSONFormatter)
	contextLogger.Level = logrus.DebugLevel
	loggerEntry = contextLogger.WithFields(logrus.Fields{"app": g_appName})

	loggerEntry.Info("Starting service" + g_appName + "...")

}

func SetAppName(app string) {
  g_appName = app
}
func Panic(args ...interface{}) {
  loggerEntry.Panic(args...)
}
func Debug(args ...interface{}) {
  loggerEntry.Debug(args...)
}
func Warn(args ...interface{}) {
  loggerEntry.Warning(args...)
}
func Error(args ...interface{}) {
  loggerEntry.Error(args...)
}
func Fatal(args ...interface{}) {
  loggerEntry.Fatal(args...)
}
func Info(args ...interface{}) {
  loggerEntry.Info(args...)
}
