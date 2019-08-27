package service

import (
	"context"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/olivere/elastic"
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
				if good.GoodsId == goods.GoodsId {
					good.TotalNum += 1
				}
			}
		}
	} else {
		currWorker := &idworker.IdWorker{}
		currWorker.InitIdWorker(1000, 1)
		newId, err := currWorker.NextId()
		if err != nil {

		}
		thiscart.CartId = newId
		thiscart.Goods[0] = goods
		thiscart.TotalNum = 1
		thiscart.TotalWeight = goods.GoodsWeight
		thiscart.TotalPrice = goods.GoodsPrice

	}

	resp.Data = thiscart
	return resp
}

//立即购买
func BuyNow(user_id int64, goods model.CartGoods) Response {

	currWorker := &idworker.IdWorker{}
	currWorker.InitIdWorker(1000, 1)
	newId, err := currWorker.NextId()
	if err != nil {

		resp.StatusCode = 400
		resp.Msg = "内部错误"
	}

	var cart model.Cart
	cart.UserId = user_id
	cart.Goods[0] = goods
	cart.TotalPrice = goods.GoodsPrice
	cart.TotalWeight = goods.GoodsWeight
	cart.TotalNum = 1
	cart.CartId = newId
	resp.Data = cart
	return resp
}

func updateNum(user_id, goods_id, cart_id int64, typ string) Response {

	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id)).Must(elastic.NewTermQuery("cart_id", cart_id))
	result, err := es.Search().Index().Type().Query(query).Size(1).Do(context.Background())
	if err != nil {

	}

	var cart model.Cart
	for _, hit := range result.Hits.Hits {
		cart := model.Cart{}
		cart.UnmarshalJSON(*hit.Source)
		for _, good := range cart.Goods {
			if good.GoodsId == goods_id {
				if typ == "raise" {
					good.TotalNum += 1
				} else if typ == "reduce" {
					good.TotalNum -= 1
				}

			}
		}
	}

	resp.Data = cart
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
			if good.GoodsId == goods_id {
				continue
			}
			cartGoods = append(cartGoods, good)
		}
		cart.Goods = cartGoods
	}

	resp.Data = cart
	return resp
}
