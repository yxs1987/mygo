package handle

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

// raw handler
func Raw(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		h(ctx)
		return
	})
}

// Write to response body
func WriteString(ctx *fasthttp.RequestCtx, result string) {
	ctx.WriteString(result)
}

func WriteBytes(ctx *fasthttp.RequestCtx, result []byte) {
	ctx.Write(result)
}

func CommonWriteSuccess(ctx *fasthttp.RequestCtx, resp Response) {
	resp.StatusCode = 200
	resp.Msg = "ok"
	CommonWrite(ctx, resp)
}

func CommonWriteError(ctx *fasthttp.RequestCtx, resp Response) {
	resp.StatusCode = 400
	resp.Msg = "error"
	CommonWrite(ctx, resp)
}

func CommonWrite(ctx *fasthttp.RequestCtx, resp Response) {
	resultMap := Response{
		StatusCode: resp.StatusCode,
		Msg:        resp.Msg,
		Data:       resp.Data,
	}
	result, _ := json.Marshal(resultMap)
	WriteBytes(ctx, result)
}
