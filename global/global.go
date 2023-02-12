package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/smartwalle/alipay/v3"
	"github.com/streadway/amqp"
	proto "pay/api/qvbilam/pay/v1"
	userProto "pay/api/qvbilam/user/v1"
	"pay/config"
)

var (
	Trans              ut.Translator // 表单验证
	ServerConfig       *config.ServerConfig
	MessageQueueClient *amqp.Connection
	PayServerClient    proto.PayClient
	UserServerClient   userProto.UserClient
	AlipayClient       *alipay.Client
)
