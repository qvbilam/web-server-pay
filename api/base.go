package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thinkeridea/go-extend/exnet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"pay/global"
	"strings"
)

const responseFieldData = "data"
const responseFieldMessage = "msg"
const responseFieldCode = "code"
const responseFiledErrors = "errors"
const responseFiledListTotal = "total"

// Success 成功
func Success(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		responseFieldData:    data,
		responseFieldMessage: message,
	})
}

func SuccessList(ctx *gin.Context, data interface{}, total int64) {
	ctx.JSON(http.StatusOK, gin.H{
		responseFiledListTotal: total,
		responseFieldData:      data,
	})
}

func SuccessNotMessage(ctx *gin.Context, data interface{}) {
	Success(ctx, data, "")
}

// SuccessNotContent 成功无返回内容
func SuccessNotContent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// Error 错误
func Error(ctx *gin.Context, message string) {
	ErrorCustom(ctx, message, nil, -1)
}

// ErrorCustom 自定义错误
func ErrorCustom(ctx *gin.Context, message string, errors interface{}, code int) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		responseFieldCode:    code,
		responseFieldMessage: message,
		responseFiledErrors:  errors,
	})
}

// ErrorUnauthorized 未登陆或登陆信息失效
func ErrorUnauthorized(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		responseFieldMessage: message,
	})
}

// ErrorForbidden 无权限
func ErrorForbidden(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusForbidden, gin.H{
		responseFieldMessage: message,
	})
}

func ErrorAlreadyExists(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusFound, gin.H{
		responseFieldMessage: message,
	})
}

// ErrorNotfound 没有找到
func ErrorNotfound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		responseFieldMessage: message,
	})
}

// ErrorInternal 服务器错误
func ErrorInternal(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		responseFieldMessage: message,
	})
	return
}

// ErrorUnprocessableEntity 请求格式正确,有语义错误
func ErrorUnprocessableEntity(ctx *gin.Context, errors interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		responseFiledErrors: errors,
	})
}

// HandleGrpcErrorToHttp 处理 grpc 错误转换成 http 错误
func HandleGrpcErrorToHttp(ctx *gin.Context, err error) {
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				ErrorNotfound(ctx, s.Message())
			case codes.Internal:
				ErrorInternal(ctx, s.Message())
			case codes.InvalidArgument:
				Error(ctx, s.Message())
			case codes.Unavailable:
				Error(ctx, s.Message())
			case codes.AlreadyExists:
				Error(ctx, s.Message())
			default:
				ErrorInternal(ctx, fmt.Sprintf("code: %d, message: %s", s.Code(), s.Message()))
			}

			return
		}
	}
}

// HandleValidateError 处理参数验证错误
func HandleValidateError(ctx *gin.Context, err error) {
	// 如果翻译错误
	errs, ok := err.(validator.ValidationErrors)
	if !ok { // 直接返回错误信息
		Error(ctx, err.Error())
		return
	}

	// 翻译成功后返回翻译之后的错误信息
	ErrorUnprocessableEntity(ctx, removeTopStruct(errs.Translate(global.Trans)))
	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}

	// 循环处理传进来的错误信息
	for field, err := range fields {
		// 获取第一个 . 出现的位置
		position := strings.Index(field, ".")
		//fmt.Println(field, position)

		// 获取截取后的 key 名
		key := field[position+1:]

		// 追加返回结果
		res[key] = err
	}
	return res
}

func GetClientIP(ctx *gin.Context) string {
	r := ctx.Request
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	return ip
}
