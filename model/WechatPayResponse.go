package model

//微信支付回调返回示例

//<xml>
//<appid><![CDATA[wx2421b1c4370ec43b]]></appid>
//<attach><![CDATA[支付测试]]></attach>
//<bank_type><![CDATA[CFT]]></bank_type>
//<fee_type><![CDATA[CNY]]></fee_type>
//<is_subscribe><![CDATA[Y]]></is_subscribe>
//<mch_id><![CDATA[10000100]]></mch_id>
//<nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
//<openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
//<out_trade_no><![CDATA[1409811653]]></out_trade_no>
//<result_code><![CDATA[SUCCESS]]></result_code>
//<return_code><![CDATA[SUCCESS]]></return_code>
//<sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
//<sub_mch_id><![CDATA[10000100]]></sub_mch_id>
//<time_end><![CDATA[20140903131540]]></time_end>
//<total_fee>1</total_fee>
//<coupon_fee_0><![CDATA[10]]></coupon_fee_0>
//<coupon_count><![CDATA[1]]></coupon_count>
//<coupon_type><![CDATA[CASH]]></coupon_type>
//<coupon_id><![CDATA[10000]]></coupon_id>
//<trade_type><![CDATA[JSAPI]]></trade_type>
//<transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
//</xml>

type PayResponse struct {
	AppId         string `json:"app_id"`
	Attach        string `json:"attach"`
	BankType      string `json:"bank_type"`
	FeeType       string `json:"fee_type"`
	IsSubscribe   string `json:"is_subscribe"`
	MchId         string `json:"mch_id"`
	NonceStr      string `json:"nonce_str"`
	OpenId        string `json:"open_id"`
	OutTradeNo    string `json:"out_trade_no"`
	ResultCode    string `json:"result_code"`
	ReturnCode    string `json:"return_code"`
	Sign          string `json:"sign"`
	SubMchId      string `json:"sub_mch_id"`
	TimeEnd       string `json:"time_end"`
	TotalFee      int64  `json:"total_fee"`
	TradeType     string `json:"trade_type"`
	TransactionId string `json:"transaction_id"`
}

type HttpResponse struct {
	Authorization string `json:"authorization"`
}
