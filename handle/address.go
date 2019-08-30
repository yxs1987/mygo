package handle

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"mygo/model"
	"mygo/service"
	"strconv"
)

func AddAddress(ctx *fasthttp.RequestCtx) {
	fmt.Println(12121)
	uid := ctx.Request.Header.Peek("uid")
	m := model.Address{}
	err := m.UnmarshalJSON(ctx.Request.Body())

	if err != nil {
		resp.Msg = "地址信息错误"
		CommonWriteError(ctx, resp)
		return
	}

	id, _ := strconv.ParseInt(string(uid), 10, 64)
	result := service.AddAddress(id, m)
	resp.Data = result.Data
	resp.Msg = result.Msg
	resp.StatusCode = result.StatusCode
	CommonWrite(ctx, resp)
}
