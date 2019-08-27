package model

type Cart struct {
	CartId      int64       `json:"cart_id"`
	UserId      int64       `json:"user_id"`
	Goods       []CartGoods `json:"goods"`
	TotalPrice  float64     `json:"total_price"`
	CreatedAt   string      `json:"created_at"`
	TotalNum    int         `json:"total_num"`
	TotalWeight float64     `json:"total_weight"`
}

type CartGoods struct {
	GoodsId     int64        `json:"goods_id"`
	GoodsName   string       `json:"goods_name"`
	GoodsPrice  float64      `json:"goods_price"`
	GoodsWeight float64      `json:"goods_weight"`
	GoodsImage  []GoodsImage `json:"image"`
	TotalWeight float64      `json:"total_weight"`
	TotalNum    int          `json:"total_num"`
	TotalPrice  float64      `json:"total_price"`
}
