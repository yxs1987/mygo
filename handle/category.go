package handle

import (
	"github.com/valyala/fasthttp"
	"mygo/service"
	"strconv"
)

var resp Response

func GetCategoryById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id")
	int64, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		resp.Msg = "id输入错误，请确认"
		CommonWriteError(ctx, resp)
		return
	}
	result := service.GetCategory(int64)
	resp.Data = result.Data

	CommonWriteSuccess(ctx, resp)
}
