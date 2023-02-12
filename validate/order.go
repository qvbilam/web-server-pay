package validate

type OrderValidate struct {
	DeliveryID int64  `form:"delivery_id" json:"delivery_id" binding:"omitempty"`
	GoodsType  string `form:"goods_type" json:"goods_type" binding:"required"`
	GoodsId    int64  `form:"goods_id" json:"goods_id" binding:"required"`
	Count      int64  `form:"count" json:"count" binding:"required"`
	Remark     string `form:"remark" json:"remark" binding:"omitempty"`
}
