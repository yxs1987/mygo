package service

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/uniplaces/carbon"
	"go.uber.org/zap"
	"mygo/common"
	"mygo/logging"
	"mygo/model"
)

var resp Response
var log *zap.Logger

func init() {
	resp.StatusCode = 200
	resp.Msg = "ok"
	log = logging.ZApLogger()
}

func AddGoods(user_id, cart_id int64, goods model.CartGoods) Response {
	var resp Response
	var thiscart model.Cart
	//查询是否保存过该用户的购物车信息
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id)).Must(elastic.NewTermQuery("cart_id", cart_id))
	result, err := es.Search().Index(common.ES_CART).Type(common.ES_CART).Size(1).Query(query).Do(context.Background())
	if err != nil {
		resp.Msg = ""
		resp.StatusCode = 400
		resp.Data = nil
	}
	if result.Hits.TotalHits == 1 {
		for _, hit := range result.Hits.Hits {
			cart := model.Cart{}
			cart.UnmarshalJSON(*hit.Source)

			for _, good := range cart.Goods {
				if good.Goods.GoodsId == goods.Goods.GoodsId {
					good.TotalNum += 1
					good.TotalPrice = float64(good.TotalNum) * good.Goods.GoodsPrice
				}
			}
		}
	} else {
		thiscart.CartId = common.GetNextId()
		thiscart.Goods[0] = goods
		thiscart.TotalNum = 1
		thiscart.TotalWeight = goods.Goods.GoodsWeight
		thiscart.TotalPrice = goods.Goods.GoodsPrice

	}

	resp.Data = thiscart
	return resp
}

//立即购买
func BuyNow(user_id, goods_id int64) Response {
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id)).Must(elastic.NewTermQuery("type", 2))
	_, err := es.DeleteByQuery().Index(common.ES_CART).Type(common.ES_CART).Query(query).Size(1).Do(context.Background())

	if err != nil {
		resp.StatusCode = 400
		resp.Msg = "查询错误"
		return resp
	}
	var cart model.Cart

	query2 := elastic.NewTermQuery("goods_id", goods_id)
	goodsResult, err := es.Search().Index(common.ES_GOODS).Type(common.ES_GOODS).Size(1).Query(query2).Do(context.Background())
	if err != nil {
		resp.Msg = "没有该商品"
		resp.StatusCode = 400
		return resp
	}
	if goodsResult.Hits.TotalHits == 0 {
		resp.Msg = "没有该商品"
		resp.StatusCode = 400
		return resp
	}

	var goods model.Goods
	for _, hit := range goodsResult.Hits.Hits {
		err := goods.UnmarshalJSON(*hit.Source)
		if err != nil {
			resp.Msg = "商品数据异常"
			resp.StatusCode = 400
			return resp
		}
	}

	var cartGoods model.CartGoods

	cartGoods.TotalWeight = goods.GoodsWeight
	cartGoods.TotalPrice = goods.GoodsPrice
	cartGoods.TotalNum = 1
	cartGoods.Goods = goods

	cart.UserId = user_id
	cart.Goods = append(cart.Goods, cartGoods)

	cart.TotalPrice = goods.GoodsPrice
	cart.TotalWeight = goods.GoodsWeight
	cart.TotalNum = 1
	cart.CartId = carbon.Now().Unix()
	cart.CreatedAt = carbon.Now().Format(carbon.DefaultFormat)

	_, err = es.Index().Index(common.ES_CART).Type(common.ES_CART).BodyJson(cart).Do(context.Background())
	if err != nil {
		resp.Msg = "购买失败"
		resp.StatusCode = 400
	} else {
		resp.Data = cart

	}

	return resp
}

func updateNum(user_id, goods_id, cart_id int64, typ string) Response {

	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id)).Must(elastic.NewTermQuery("cart_id", cart_id))
	result, err := es.Search().Index().Type().Query(query).Size(1).Do(context.Background())
	var cart model.Cart
	if err != nil {
		resp.Msg = "更新出错"
		resp.StatusCode = 400
	} else {
		for _, hit := range result.Hits.Hits {
			cart.UnmarshalJSON(*hit.Source)
			for _, good := range cart.Goods {
				if good.Goods.GoodsId == goods_id {
					if typ == "raise" {
						good.TotalNum += 1
					} else if typ == "reduce" {
						good.TotalNum -= 1
					}

				}
			}
		}
		resp.Data = cart
	}

	return resp
}

//增加数量
func RaiseNum(user_id, goods_id, cart_id int64) Response {
	resp.Data = updateNum(user_id, goods_id, cart_id, "raise")
	return resp
}

//减少数量
func ReduceNum(user_id, goods_id, cart_id int64) Response {
	resp.Data = updateNum(user_id, goods_id, cart_id, "reduce")
	return resp
}

//删除购物车中商品
func DeleteGoods(user_id, goods_id, cart_id int64) Response {
	var cart model.Cart
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id)).Must(elastic.NewTermQuery("cart_id", cart_id)).Must(elastic.NewNestedQuery("goods", elastic.NewTermQuery("goods.goods_id", goods_id)))
	result, err := es.Search().Index().Type().Size(1).Query(query).Do(context.Background())
	if err != nil {

	}
	if result.Hits.TotalHits == 0 {
		resp.StatusCode = 400
		resp.Msg = ""
	} else {
		var cartGoods []model.CartGoods
		for _, good := range cart.Goods {
			if good.Goods.GoodsId == goods_id {
				continue
			}
			cartGoods = append(cartGoods, good)
		}
		cart.Goods = cartGoods
	}

	resp.Data = cart
	return resp
}
