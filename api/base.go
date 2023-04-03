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
const responseFiledList = "list"

// Success 成功
func Success(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		responseFieldData:    data,
		responseFieldMessage: message,
	})
}

func SuccessList(ctx *gin.Context, data interface{}, total int64) {
	Success(ctx, gin.H{
		responseFiledListTotal: total,
		responseFiledList:      data,
	}, "")
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
	ErrorCustom(ctx, http.StatusBadRequest, message)
}

// ErrorCustom 自定义错误
func ErrorCustom(ctx *gin.Context, httpCode int, message string) {
	ctx.JSON(httpCode, gin.H{responseFieldMessage: message})
}

// ErrorAlreadyExists 资源存在 302
func ErrorAlreadyExists(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusFound, gin.H{
		responseFieldMessage: message,
		responseFieldData:    data,
	})
}

// ErrorUnauthorized 未登陆或登陆信息失效 401
func ErrorUnauthorized(ctx *gin.Context, message string) {
	ErrorCustom(ctx, http.StatusUnauthorized, message)
}

// ErrorPayment 需要支付 402
func ErrorPayment(ctx *gin.Context, message string) {
	ErrorCustom(ctx, http.StatusPaymentRequired, message)
}

// ErrorForbidden 无权限 403
func ErrorForbidden(ctx *gin.Context, message string) {
	ErrorCustom(ctx, http.StatusForbidden, message)
}

// ErrorNotfound 没有找到 404
func ErrorNotfound(ctx *gin.Context, message string) {
	ErrorCustom(ctx, http.StatusNotFound, message)
}

func ErrorTimeout(ctx *gin.Context) {
	ctx.JSON(http.StatusRequestTimeout, gin.H{
		responseFieldMessage: "请求超时",
	})
}

// ErrorUnprocessableEntity 请求格式正确,有语义错误 422
func ErrorUnprocessableEntity(ctx *gin.Context, errors interface{}) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		responseFieldMessage: "参数错误",
		responseFiledErrors:  errors,
	})
}

func ErrorTooManyRequests(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusTooManyRequests, gin.H{
		responseFieldMessage: "请求频繁,请稍后再试",
	})
}

// ErrorInternal 服务器错误 500
func ErrorInternal(ctx *gin.Context, message string) {
	ErrorCustom(ctx, http.StatusInternalServerError, message)
}

// HandleGrpcErrorToHttp 处理 grpc 错误转换成 http 错误
func HandleGrpcErrorToHttp(ctx *gin.Context, err error) {
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.DeadlineExceeded:
				ErrorTimeout(ctx)
			case codes.NotFound:
				ErrorNotfound(ctx, s.Message())
			case codes.Internal:
				ErrorInternal(ctx, s.Message())
			case codes.Unauthenticated:
				ErrorUnauthorized(ctx, s.Message())
			case codes.InvalidArgument:
				Error(ctx, s.Message())
			case codes.Unavailable:
				Error(ctx, s.Message())
			case codes.AlreadyExists:
				Error(ctx, s.Message())
			case codes.ResourceExhausted:
				ErrorTooManyRequests(ctx, s.Message())
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
