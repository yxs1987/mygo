package handle

import (
	"github.com/valyala/fasthttp"
	"mygo/model"
	"mygo/service"
	"strconv"
)

func OrderList(ctx *fasthttp.RequestCtx) {

	var order model.OrderList
	req := ctx.Request.Body()
	order.UnmarshalJSON(req)
	result := service.GetOrderList(order)
	resp.Data = result
	CommonWriteSuccess(ctx, resp)
}

//生成订单
func OrderDo(ctx *fasthttp.RequestCtx) {
	var order model.CreateOrder
	req := ctx.Request.Body()

	err := order.UnmarshalJSON(req)

	if err != nil {
		resp.Msg = "参数错误"
		CommonWriteError(ctx, resp)
	}

	user_id := ctx.Request.Header.Peek("uid")

	uid, _ := strconv.ParseInt(string(user_id), 10, 64)
	result := service.CreateOrder(order.CartId, order.AddressId, uid)
	resp.StatusCode = result.StatusCode
	resp.Msg = result.Msg
	resp.Data = result.Data
	CommonWrite(ctx, resp)
}
