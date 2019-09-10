package handle

import (
	"github.com/valyala/fasthttp"
	"mygo/service"
	"strconv"
)

func AddToCart(ctx *fasthttp.RequestCtx) {

}

func BuyNow(ctx *fasthttp.RequestCtx) {
	goods_id := ctx.QueryArgs().Peek("goods_id")
	user_id := ctx.Request.Header.Peek("uid")
	if goods_id == nil {
		resp.Msg = "缺少商品"
		CommonWriteError(ctx, resp)
		return
	}

	gid, _ := strconv.ParseInt(string(goods_id), 10, 64)
	uid, _ := strconv.ParseInt(string(user_id), 10, 64)
	result := service.BuyNow(uid, gid)
	resp.StatusCode = result.StatusCode
	resp.Msg = result.Msg
	resp.Data = result.Data
	CommonWrite(ctx, resp)
}

func RaiseNum(ctx *fasthttp.RequestCtx) {
	goods_id := ctx.UserValue("goods_id")
	user_id := ctx.Request.Header.Peek("user_id")
	cart_id := ctx.UserValue("cart_id")

	gid, _ := strconv.ParseInt(goods_id.(string), 10, 64)
	uid, _ := strconv.ParseInt(string(user_id), 10, 64)
	cid, _ := strconv.ParseInt(cart_id.(string), 10, 64)

	result := service.RaiseNum(uid, gid, cid)
	resp.StatusCode = result.StatusCode
	resp.Msg = result.Msg
	resp.Data = result.Data
	CommonWrite(ctx, resp)
}
