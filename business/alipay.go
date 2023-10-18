package business

import (
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"pay/global"
	"strconv"
)

const kSandboxURL = "https://openapi-sandbox.dl.alipaydev.com/gateway.do"

type AliPayBusiness struct {
	OrderSn string
	Subject string
	Amount  float64
}

func (b *AliPayBusiness) Web() (string, error) {
	var p = alipay.TradePagePay{}
	p.NotifyURL = global.ServerConfig.AliPayConfig.NotifyUrl // 支付通知地址
	p.ReturnURL = global.ServerConfig.AliPayConfig.ReturnUrl // 浏览器跳转页面地址
	p.Subject = b.Subject
	p.OutTradeNo = b.OrderSn
	p.TotalAmount = strconv.FormatFloat(b.Amount, 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY" // 网页支付使用: FAST_INSTANT_TRADE_PAY ;其他类型支付使用: QUICK_WAP_WAY
	url, err := global.AlipayClient.TradePagePay(p)
	if err != nil {
		zap.S().Errorw("创建支付订单失败")
		return "", err
	}
	// 替换新版沙箱环境网关
	if !global.ServerConfig.AliPayConfig.IsProduction {
		return kSandboxURL + "?" + url.RawQuery, nil
	}
	return url.String(), nil
}
