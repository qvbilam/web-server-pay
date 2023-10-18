package business

import (
	"context"
	proto "pay/api/qvbilam/pay/v1"
	"pay/global"
)

type OrderBusiness struct {
	ID         int64
	UserID     int64
	OrderSn    string
	PayType    string
	ClientType string
}

func (b *OrderBusiness) Apply() (*proto.OrderResponse, error) {
	res, err := global.PayServerClient.ApplyOrder(context.Background(), &proto.ApplyOrderRequest{
		OrderSn:    b.OrderSn,
		PayType:    b.PayType,
		ClientType: b.ClientType,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
