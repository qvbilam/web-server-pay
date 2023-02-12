package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pay/api"
	"pay/utils"
	"strconv"
	"time"
)

// LoggerToFile 日志的中间件
func LoggerToFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//fmt.Println("收到请求")
		logger, err := newLogFile("request")
		defer func(logger *zap.Logger) {
			_ = logger.Sync()
		}(logger) // 刷新缓存
		if err != nil {
			api.HandleValidateError(ctx, err)
			ctx.Abort()
			return
		}

		// 开始的时间
		startTime := time.Now()

		// 处理请求
		ctx.Next()
		// 处理请求
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)

		reqMethod := ctx.Request.Method

		// 请求路由
		reqUri := ctx.Request.RequestURI

		// 状态码
		statusCode := ctx.Writer.Status()

		// 请求IP
		clientIp := ctx.ClientIP()

		// 请求参数
		ps, _ := utils.RequestParams(ctx)

		// 请求头
		hs, _ := utils.RequestHeaders(ctx)

		params, _ := json.Marshal(ps)
		headers, _ := json.Marshal(hs)
		// 日志的格式
		logger.Info(reqUri,
			zap.Int("httpStatus", statusCode),
			zap.String("responseMilliseconds", strconv.FormatInt(latencyTime.Milliseconds(), 10)), // 毫秒
			zap.String("clientIP", clientIp),
			zap.String("requestMethod", reqMethod),
			zap.String("requestUri", reqUri),
			zap.String("requestParams", string(params)),
			zap.String("requestHeaders", string(headers)),
		)

		//logger.Info("| %3d | %13v | %15s | %s | %s | %s | %+v ",
		//	statusCode,
		//	latencyTime,
		//	clientIp,
		//	reqMethod,
		//	reqUri,
		//	clientIp,
		//	params,
		//)
	}
}

func newLogFile(name string) (*zap.Logger, error) {
	t := time.Now().Local().Format("2006-01-02")
	file := fmt.Sprintf("./logs/%s-%s.log", name, t)

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		file, // 输出到文件中, 需要注意使用的目录. 建议go build 执行
		//"stdout", // 输出到控制台中
	}
	return config.Build()
}
