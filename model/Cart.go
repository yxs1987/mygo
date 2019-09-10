package model

type Cart struct {
	CartId      int64       `json:"cart_id"`
	UserId      int64       `json:"user_id"`
	Goods       []CartGoods `json:"goods"`
	TotalPrice  float64     `json:"total_price"`
	CreatedAt   string      `json:"created_at"`
	TotalNum    int         `json:"total_num"`
	TotalWeight float64     `json:"total_weight"`
	Type        int         `json:"type"` //购物车类型 1=普通 2=立即购买
}

type CartGoods struct {
	Goods       Goods   `json:"goods"`
	TotalWeight float64 `json:"total_weight"`
	TotalNum    int     `json:"total_num"`
	TotalPrice  float64 `json:"total_price"`
}
