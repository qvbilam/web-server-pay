package ali

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"pay/api"
	proto "pay/api/qvbilam/pay/v1"
	"pay/business"
	"pay/enum"
	"pay/global"
	"pay/resource"
	"pay/validate"
	"strconv"
	"time"
)

func Web(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.OrderApplyValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	b := business.OrderBusiness{
		UserID:     userID,
		OrderSn:    request.OrderSn,
		PayType:    enum.PayTypeAlipay,
		ClientType: enum.ClientTypeWeb,
	}

	order, err := b.Apply()
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	// 生成url
	ab := business.AliPayBusiness{
		OrderSn: order.OrderSn,
		Subject: order.Subject,
		Amount:  float64(order.Amount),
	}
	fmt.Printf("ab: %+v\n", ab)

	url, err := ab.Web()
	if err != nil {
		api.Error(ctx, "生成订单失败")
		return
	}
	fmt.Println(url)
	res := resource.AliWebResource{}
	api.SuccessNotMessage(ctx, res.Resource(url, order))
}

func Notify(ctx *gin.Context) {
	// 验证签名返回参数
	notify, err := global.AlipayClient.GetTradeNotification(ctx.Request)
	if err != nil {
		api.Error(ctx, err.Error())
		return
	}

	// 更新交易状态
	status := enum.PayStatusWait
	if notify.TradeStatus == alipay.TradeStatusSuccess {
		status = enum.PayStatusSuccess
	}
	if notify.TradeStatus == alipay.TradeStatusFinished {
		status = enum.PayStatusFinished
	}
	if notify.TradeStatus == alipay.TradeStatusClosed {
		status = enum.PayStatusClosed
	}

	payResult, _ := json.Marshal(notify)
	payAmount, _ := strconv.ParseFloat(notify.BuyerPayAmount, 64)
	// 日期转换时间
	tm, _ := time.Parse("2006-01-02 15:04:05", notify.GmtPayment)
	if _, err := global.PayServerClient.UpdateOrder(context.Background(), &proto.UpdateOrderRequest{
		OrderSn:   notify.OutTradeNo,
		TradeNo:   notify.TradeNo,
		Status:    status,
		PayAmount: float32(payAmount),
		PayResult: string(payResult),
		PayTime:   tm.Unix(),
	}); err != nil {
		api.Error(ctx, err.Error())
		return
	}
	api.Success(ctx, nil, "ok")
}
