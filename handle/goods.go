package handle

import (
	"github.com/valyala/fasthttp"
	"mygo/service"
	"strconv"
)

func GetGoodsById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id")

	gid, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		resp.Msg = "id输入错误，请确认"
		CommonWriteError(ctx, resp)
		return
	}
	resp.Data = service.GetGoodsById(gid)
	CommonWriteSuccess(ctx, resp)
}
