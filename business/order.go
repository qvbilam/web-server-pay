package business

import (
	"context"
	proto "pay/api/qvbilam/pay/v1"
	"pay/global"
)

type OrderBusiness struct {
	UserID     int64
	DeliveryID int64
	GoodsType  string
	GoodsId    int64
	Count      int64
	PayType    string
	ClientType string
	Remark     string
}

func (b *OrderBusiness) Create() (*proto.OrderResponse, error) {
	order, err := global.PayServerClient.CreateOrder(context.Background(), &proto.CreateOrderRequest{
		UserId:     b.UserID,
		GoodsType:  b.GoodsType,
		GoodsId:    b.GoodsId,
		Count:      b.Count,
		DeliveryId: b.DeliveryID,
		PayType:    b.PayType,
		ClientType: b.ClientType,
		Remark:     b.Remark,
	})
	if err != nil {
		return nil, err
	}

	return order, nil
}
