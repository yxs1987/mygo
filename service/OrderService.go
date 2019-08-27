package service

import (
	"context"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/olivere/elastic"
	"mygo/model"
)

func CreateOrder(cart model.Cart, address model.Address) Response {
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
	resp.Data = order
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
