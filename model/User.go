package model

type WechatResponse struct {
	Code     string   `json:"code"`
	UserInfo UserInfo `json:"user_info"`
}

type Tm struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserInfo struct {
	Tm
	UserId          int64     `json:"user_id"`
	OpenId          string    `json:"open_id"`
	Nickname        string    `json:"nickName"`
	Gender          int       `json:"gender"`
	Language        string    `json:"language"`
	City            string    `json:"city"`
	Province        string    `json:"province"`
	Country         string    `json:"country"`
	AvatarUrl       string    `json:"avatar_url"`
	Address         []Address `json:"address"`
	LatestLoginTime string    `json:"latest_login_time"`
}

type Address struct {
	Tm
	AddressId string `json:"address_id"`
	Name      string `json:"name"`
	UserId    int64  `json:"user_id"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	Detail    string `json:"detail"`
	Mobile    int64  `json:"mobile"`
	Status    int    `json:"status"`
}
