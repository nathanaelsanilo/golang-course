package logger

import (
	"go.uber.org/zap"
)

func BuildSugarLogger() *zap.SugaredLogger {
	log, _ := zap.NewProduction()
	sugar := log.Sugar()
	return sugar
}

// ... variadic params
func BuildInfo(template string, args ...interface{}) {
	sugar := BuildSugarLogger()
	sugar.Infof(template, args...)
}
