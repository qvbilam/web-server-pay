package initialize

import (
	"go.uber.org/zap"
)

func InitLogger() {
	// 定义全局日志
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}
