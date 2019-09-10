package handle

import (
	"bufio"
	"fmt"
	"github.com/valyala/fasthttp"
	"mygo/model"
	"mygo/service"
)

var wx model.WechatResponse

func Notify(ctx *fasthttp.RequestCtx) {

	fmt.Println(12121)
	bf := &bufio.Writer{}
	bf.Write([]byte{'1', '2', '3'})
	ctx.Response.Write(bf)
	fmt.Println(ctx.Request.Body())
}

func Pay(ctx *fasthttp.RequestCtx) {

	data := ctx.Request.Body()
	var pay model.PayOrder

	err := pay.UnmarshalJSON(data)
	if err != nil {
		resp.Msg = "参数错误"
		CommonWriteError(ctx, resp)
	}
	order := service.GetOrder(pay.OrderId, pay.UserId)
	if order.OrderId == 0 {
		CommonWriteError(ctx, resp)
	}

}
