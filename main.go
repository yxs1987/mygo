package main

import (
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"mygo/handle"
	req "mygo/jwt_request"
	"mygo/logging"
	"reflect"
)

var (
	corsAllowHeaders     = "authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
	contentType          = "application/json"
)

var R struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	var r = R
	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		ctx.Response.Header.SetContentType(contentType)

		//info := ctx.Request.Header.Peek(fasthttp.HeaderAuthorization)

		token, err := req.ParseFromRequest(ctx, req.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("abc"), nil
			})

		if err == nil {
			if token.Valid {
				claims := token.Claims
				uid := GetIdFromClaims("uid", claims)
				ctx.Request.Header.Set("uid", uid)
			} else {
				r.Code = fasthttp.StatusUnauthorized
				r.Msg = "token not valid"
				bt, _ := json.Marshal(r)
				ctx.Write(bt)
				return
			}
		} else {
			r.Code = fasthttp.StatusUnauthorized
			r.Msg = fasthttp.HeaderAuthorization
			bt, _ := json.Marshal(r)
			ctx.Write(bt)
			return
		}

		next(ctx)
	}
}

// 示例 ：GetIdFromClaims("username", token.claims) 其中token是已经解密的token
func GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
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
	//api.POST("/auth/*filepath", CORS(handle.Notify))
	//api.GET("/auth/*filepath", CORS(handle.Notify))
	api.POST("/auth/category/:id", CORS(handle.GetCategoryById))
	api.GET("/auth/goods/:id", CORS(handle.GetGoodsById))
	api.POST("/auth/address/add", CORS(handle.AddAddress))
	api.GET("/auth/cart/add", CORS(handle.AddAddress))
	api.GET("/auth/cart/buyNow", CORS(handle.BuyNow))
	api.POST("/auth/order/create", CORS(handle.OrderDo))
	api.POST("/auth/order/list", CORS(handle.OrderList))

	err := fasthttp.ListenAndServe(":8080", api.Handler)

	fmt.Println("listen:8080")
	if err != nil {
		panic(err)
	}

}
