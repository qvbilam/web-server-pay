package validate

type OrderApplyValidate struct {
	OrderSn string `form:"order_sn" json:"order_sn" binding:"required"`
}
