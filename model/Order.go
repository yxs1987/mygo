package model

type Order struct {
	UserId     int64       `json:"user_id"`
	OrderGoods []CartGoods `json:"goods"`
	TotalPrice float64     `json:"total_price"`
	Consignee  string      `json:"consignee"`
	Mobile     int64       `json:"mobile"`
	Province   string      `json:"province"`
	City       string      `json:"city"`
	District   string      `json:"district"`
	Detail     string      `json:"detail"`
	OrderId    int64       `json:"order_id"`
	PayStatus  int         `json:"pay_status"`
	PayPrice   float64     `json:"pay_price"`
	CreatedAt  string      `json:"create_at"`

	UpdatedAt     string `json:"updated_at"`
	NonceStr      string `json:"nonce_str"`
	SignType      string `json:"sign_type"`
	Openid        string `json:"openid"`
	IsSubscribe   string `json:"is_subscribe"`
	TradeType     string `json:"trade_type"`
	BankType      string `json:"bank_type"`
	TransactionId string `json:"transaction_id"`
	PayTimeEnd    string `json:"pay_time_end"`
}
