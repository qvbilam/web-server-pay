package initialize

import (
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"pay/global"
)

func InitAliPay() {
	appId := global.ServerConfig.AliPayConfig.AppID
	privateKey := global.ServerConfig.AliPayConfig.PrivateKey
	isProduction := global.ServerConfig.AliPayConfig.IsProduction
	aliPublicKey := global.ServerConfig.AliPayConfig.AliPublicKey
	var client, err = alipay.New(appId, privateKey, isProduction) // false 使用沙箱环境
	if err != nil {
		zap.S().Fatalf("支付初始化失败")
	}
	global.AlipayClient = client
	if err = global.AlipayClient.LoadAliPayPublicKey(aliPublicKey); err != nil {
		zap.S().Fatalf("支付公钥加载失败")
	}
}
