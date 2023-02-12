package resource

import proto "pay/api/qvbilam/pay/v1"

type AliWebResource struct {
	OrderSn string  `json:"order_sn"`
	Subject string  `json:"subject"`
	Amount  float64 `json:"amount"`
	Url     string  `json:"url"`
}

func (s *AliWebResource) Resource(url string, p *proto.OrderResponse) *AliWebResource {
	return &AliWebResource{
		OrderSn: p.OrderSn,
		Subject: p.Subject,
		Amount:  float64(p.Amount),
		Url:     url,
	}
}
