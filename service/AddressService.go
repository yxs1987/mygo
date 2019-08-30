package service

import (
	"context"
	"github.com/olivere/elastic"
	"go.uber.org/zap"
	"mygo/common"
	"mygo/model"
)

func AddAddress(user_id int64, address model.Address) Response {
	var addr []model.Address
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id)).MustNot(elastic.NewTermQuery("status", 13))
	cnt, err := es.Count().Index(common.ES_ADDRESS).Type(common.ES_ADDRESS).Query(query).Do(context.Background())
	if err != nil {
		resp.StatusCode = 400
		resp.Msg = err.Error()
		return resp
	}
	if cnt >= 10 {
		log.Info("地址保存数量超过限制", zap.Int64("user_id", user_id), zap.Any("address", address))
		resp.StatusCode = 400
		resp.Msg = ""
	} else {

		address.UserId = user_id
		address.AddressId = common.GetNextId()

		_, err := es.Index().Index(common.ES_ADDRESS).Type(common.ES_ADDRESS).BodyJson(address).Do(context.Background())
		if err != nil {
			resp.Msg = err.Error()
			resp.StatusCode = 400
		}

		result, _ := es.Search().Index().Query(query).Do(context.Background())
		for _, hit := range result.Hits.Hits {
			var add model.Address
			add.UnmarshalJSON(*hit.Source)
			addr = append(addr, add)
		}

		user, err := es.Search().Index(common.ES_USER).Query(elastic.NewTermQuery("user_id", user_id)).Size(1).Do(context.Background())
		if err == nil {
			for _, hit := range user.Hits.Hits {
				up := map[string]interface{}{"address": addr}
				es.Update().Index(common.ES_USER).Type(common.ES_USER).Id(hit.Id).Doc(up).Do(context.Background())
			}
		}

	}

	resp.Data = address
	return resp
}

func DelAddress(user_id, address_id int64) Response {
	var adds []model.Address
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id)).Must(elastic.NewTermQuery("address_id", address_id)).MustNot(elastic.NewTermQuery("status", 13))
	address, _ := es.Search().Index().Type().Size(1).Query(query).Do(context.Background())
	if address.Hits.TotalHits != 1 {
		resp.Msg = ""
		resp.StatusCode = 400
	}
	query = elastic.NewBoolQuery().Must(elastic.NewTermQuery("user_id", user_id)).MustNot(elastic.NewTermQuery("address_id", address_id)).MustNot(elastic.NewTermQuery("status", 13))
	result, _ := es.Search().Index().Query(query).Do(context.Background())
	for _, hit := range result.Hits.Hits {
		c := model.Address{}
		c.UnmarshalJSON(*hit.Source)
		adds = append(adds, c)
	}
	//删除地址列表
	go func() {
		for _, hit := range address.Hits.Hits {
			up := map[string]int{"status": 13}
			es.Update().Id(hit.Id).Index(hit.Index).Type(hit.Type).Doc(up).Do(context.Background())
		}
	}()
	//更新用户地址列表属性
	go func() {
		user, _ := es.Search().Index().Size(1).Query(elastic.NewTermQuery("user_id", user_id)).Do(context.Background())
		for _, hit := range user.Hits.Hits {
			up := map[string]interface{}{"address": adds}
			es.Update().Id(hit.Id).Index(common.ES_USER).Type(common.ES_USER).Doc(up).Do(context.Background())
		}
	}()

	resp.Data = adds
	return resp
}
