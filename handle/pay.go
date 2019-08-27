package handle

import (
	"bufio"
	"fmt"
	"github.com/valyala/fasthttp"
	"mygo/model"
)

var wx model.WechatResponse

func Notify(ctx *fasthttp.RequestCtx) {

	fmt.Println(12121)
	bf := &bufio.Writer{}
	bf.Write([]byte{'1', '2', '3'})
	ctx.Response.Write(bf)
	fmt.Println(ctx.Request.Body())
}
