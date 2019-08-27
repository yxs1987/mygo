package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"mygo/handle"
	"mygo/logging"
	"mygo/model"
)

var (
	corsAllowHeaders     = "authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
	contentType          = "application/json"
)

var hp model.HttpResponse

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		ctx.Response.Header.SetContentType(contentType)

		info := ctx.Request.Header.Peek(fasthttp.HeaderAuthorization)
		fmt.Println(string(info))

		next(ctx)
	}
}

var ZapLog *zap.Logger

func init() {
	ZapLog = logging.ZApLogger()
}

func main() {
	api := fasthttprouter.New()
	api.POST("/api/user/login", handle.Login)

	//api.GET("/good", handle.GoodList)
	//api.GET("/good/get/:id", handle.GoodView)
	//api.GET("/category",handle.CategoryList)
	//api.GET("/category/:id",handle.CategoryGoods)
	//api.POST("/cart/add",handle.AddToCart)

	//api.POST("/".)
	api.GET("/test", handle.Notify)

	err := fasthttp.ListenAndServe(":8080", CORS(api.Handler))
	fmt.Println("listen:8080")
	if err != nil {
		panic(err)
	}

}
