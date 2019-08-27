package model

type Goods struct {
	GoodsId     int64        `json:"goods_id"`
	GoodsName   string       `json:"goods_name"`
	GoodsPrice  float64      `json:"goods_price"`
	GoodsWeight float64      `json:"goods_weight"`
	Image       []GoodsImage `json:"image"`
	Content     string       `json:"content"`
}

type GoodsImage struct {
	ImageId int64  `json:"image_id"`
	Url     string `json:"url"`
	GoodsId int64  `json:"goods_id"`
	Type    string `json:"type"` //图片类型，商品细节图=1，购物车小图=2
}
