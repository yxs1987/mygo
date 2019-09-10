package service

import (
	"context"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/olivere/elastic"
	"math"
	"mygo/common"
	"mygo/model"
)

func CreateOrder(cart_id, address_id, user_id int64) Response {
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("cart_id", cart_id)).Must(elastic.NewTermQuery("user_id", user_id))
	cartresult, _ := es.Search().Index(common.ES_CART).Type(common.ES_CART).Size(1).Query(query).Do(context.Background())
	if cartresult.Hits.TotalHits == 0 {
		resp.Msg = "没有购物信息"
		resp.StatusCode = 400
		return resp
	}

	var cart model.Cart
	var address model.Address
	for _, hit := range cartresult.Hits.Hits {

		cart.UnmarshalJSON(*hit.Source)
	}
	query2 := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id)).Must(elastic.NewTermQuery("address_id", address_id)).Must(elastic.NewTermQuery("status", 1))

	addresult, _ := es.Search().Index(common.ES_ADDRESS).Type(common.ES_ADDRESS).Query(query2).Size(1).Do(context.Background())
	if addresult.Hits.TotalHits == 0 {
		resp.Msg = "地址信息错误"
		resp.StatusCode = 400
		return resp
	}
	for _, hit := range addresult.Hits.Hits {
		address.UnmarshalJSON(*hit.Source)
	}

	var order model.Order
	order.CreatedAt = ""
	order.UpdatedAt = ""

	currWorker := &idworker.IdWorker{}
	currWorker.InitIdWorker(1000, 1)
	newId, err := currWorker.NextId()
	if err != nil {
		resp.Msg = err.Error()
		return resp
	}

	order.OrderId = newId
	order.TotalPrice = cart.TotalPrice
	order.OrderGoods = cart.Goods
	order.City = address.City
	order.Province = address.Province
	order.District = address.District
	order.UserId = cart.UserId
	order.Consignee = address.Name
	order.Mobile = address.Mobile
	order.Detail = address.Detail
	_, err = es.Index().Index(common.ES_ORDER).Type(common.ES_ORDER).BodyJson(order).Do(context.Background())
	if err != nil {
		resp.StatusCode = 400
		resp.Msg = "创建订单失败"
	} else {
		resp.Data = order
		resp.StatusCode = 200
	}

	return resp
}

func GetOrderByStatus(user_id int64, pay_status int) Response {
	var orders []model.Order

	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id))
	if pay_status != 0 {
		query.Must(elastic.NewTermQuery("pay_status", pay_status))
	}
	result, _ := es.Search().Index().Type().Query(query).Size(1).Do(context.Background())
	for _, hit := range result.Hits.Hits {
		var order model.Order
		order.UnmarshalJSON(*hit.Source)
		orders = append(orders, order)
	}
	resp.Data = orders
	return resp
}

func UpdateOrder() {

}

func GetOrder(order_id, user_id int64) model.Order {
	var order model.Order
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("order_id", order_id)).Must(elastic.NewTermQuery("user_id", user_id))
	result, err := es.Search().Index(common.ES_ORDER).Type(common.ES_ORDER).Size(1).Query(query).Do(context.Background())
	if err != nil {
		return order
	}
	for _, hit := range result.Hits.Hits {
		order.UnmarshalJSON(*hit.Source)
	}
	return order
}

func GetOrderList(m model.OrderList) interface{} {
	var o []model.Order
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", m.UserId))
	if m.Page.PageNum == 0 {
		m.Page.PageNum = 1
	}
	if m.Page.PageSize == 0 {
		m.Page.PageSize = 20
	}
	if m.OrderStatus != 0 {
		query.Must(elastic.NewTermQuery("order_status", m.OrderStatus))
	}
	if m.PayStatus != 0 {
		query.Must(elastic.NewTermQuery("pay_status", m.PayStatus))
	}
	result, err := es.Search().
		Index(common.ES_ORDER).
		Type(common.ES_ORDER).
		Query(query).
		From(m.Page.PageNum).
		Size(m.Page.PageSize).
		Do(context.Background())

	resArr := map[string]interface{}{
		"total":       0,
		"currentPage": 0,
		"totalPage":   0,
		"pageSize":    0,
		"list":        []interface{}{},
	}

	if err != nil {
		return resArr
	}
	for _, hit := range result.Hits.Hits {
		order := model.Order{}
		order.UnmarshalJSON(*hit.Source)
		o = append(o, order)
	}
	resArr["total"] = result.Hits.TotalHits
	resArr["currentPage"] = m.Page.PageNum
	resArr["totalPage"] = int(math.Ceil(float64(result.Hits.TotalHits) / float64(m.Page.PageSize)))
	resArr["pageSize"] = m.Page.PageSize
	resArr["total"] = result.Hits.TotalHits
	resArr["list"] = o
	return resArr
}
