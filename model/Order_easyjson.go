// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson4f2ebbc2DecodeMygoModel(in *jlexer.Lexer, out *Order) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.UserId = int64(in.Int64())
		case "goods":
			if in.IsNull() {
				in.Skip()
				out.OrderGoods = nil
			} else {
				in.Delim('[')
				if out.OrderGoods == nil {
					if !in.IsDelim(']') {
						out.OrderGoods = make([]CartGoods, 0, 1)
					} else {
						out.OrderGoods = []CartGoods{}
					}
				} else {
					out.OrderGoods = (out.OrderGoods)[:0]
				}
				for !in.IsDelim(']') {
					var v1 CartGoods
					(v1).UnmarshalEasyJSON(in)
					out.OrderGoods = append(out.OrderGoods, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "total_price":
			out.TotalPrice = float64(in.Float64())
		case "consignee":
			out.Consignee = string(in.String())
		case "mobile":
			out.Mobile = int64(in.Int64())
		case "province":
			out.Province = string(in.String())
		case "city":
			out.City = string(in.String())
		case "district":
			out.District = string(in.String())
		case "detail":
			out.Detail = string(in.String())
		case "order_id":
			out.OrderId = int64(in.Int64())
		case "pay_status":
			out.PayStatus = int(in.Int())
		case "pay_price":
			out.PayPrice = float64(in.Float64())
		case "create_at":
			out.CreatedAt = string(in.String())
		case "updated_at":
			out.UpdatedAt = string(in.String())
		case "nonce_str":
			out.NonceStr = string(in.String())
		case "sign_type":
			out.SignType = string(in.String())
		case "openid":
			out.Openid = string(in.String())
		case "is_subscribe":
			out.IsSubscribe = string(in.String())
		case "trade_type":
			out.TradeType = string(in.String())
		case "bank_type":
			out.BankType = string(in.String())
		case "transaction_id":
			out.TransactionId = string(in.String())
		case "pay_time_end":
			out.PayTimeEnd = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4f2ebbc2EncodeMygoModel(out *jwriter.Writer, in Order) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.UserId))
	}
	{
		const prefix string = ",\"goods\":"
		out.RawString(prefix)
		if in.OrderGoods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.OrderGoods {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"total_price\":"
		out.RawString(prefix)
		out.Float64(float64(in.TotalPrice))
	}
	{
		const prefix string = ",\"consignee\":"
		out.RawString(prefix)
		out.String(string(in.Consignee))
	}
	{
		const prefix string = ",\"mobile\":"
		out.RawString(prefix)
		out.Int64(int64(in.Mobile))
	}
	{
		const prefix string = ",\"province\":"
		out.RawString(prefix)
		out.String(string(in.Province))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"district\":"
		out.RawString(prefix)
		out.String(string(in.District))
	}
	{
		const prefix string = ",\"detail\":"
		out.RawString(prefix)
		out.String(string(in.Detail))
	}
	{
		const prefix string = ",\"order_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.OrderId))
	}
	{
		const prefix string = ",\"pay_status\":"
		out.RawString(prefix)
		out.Int(int(in.PayStatus))
	}
	{
		const prefix string = ",\"pay_price\":"
		out.RawString(prefix)
		out.Float64(float64(in.PayPrice))
	}
	{
		const prefix string = ",\"create_at\":"
		out.RawString(prefix)
		out.String(string(in.CreatedAt))
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.String(string(in.UpdatedAt))
	}
	{
		const prefix string = ",\"nonce_str\":"
		out.RawString(prefix)
		out.String(string(in.NonceStr))
	}
	{
		const prefix string = ",\"sign_type\":"
		out.RawString(prefix)
		out.String(string(in.SignType))
	}
	{
		const prefix string = ",\"openid\":"
		out.RawString(prefix)
		out.String(string(in.Openid))
	}
	{
		const prefix string = ",\"is_subscribe\":"
		out.RawString(prefix)
		out.String(string(in.IsSubscribe))
	}
	{
		const prefix string = ",\"trade_type\":"
		out.RawString(prefix)
		out.String(string(in.TradeType))
	}
	{
		const prefix string = ",\"bank_type\":"
		out.RawString(prefix)
		out.String(string(in.BankType))
	}
	{
		const prefix string = ",\"transaction_id\":"
		out.RawString(prefix)
		out.String(string(in.TransactionId))
	}
	{
		const prefix string = ",\"pay_time_end\":"
		out.RawString(prefix)
		out.String(string(in.PayTimeEnd))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Order) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f2ebbc2EncodeMygoModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Order) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f2ebbc2EncodeMygoModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Order) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f2ebbc2DecodeMygoModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Order) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f2ebbc2DecodeMygoModel(l, v)
}
