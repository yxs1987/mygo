package service

import (
	"context"
	"github.com/olivere/elastic"
	"mygo/model"
)

func init() {

}

func GetGoodsById(id int64) model.Goods {
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("goods_id", id)).Must(elastic.NewTermQuery("status", 1))
	result, _ := es.Search().Index().Type().Query(query).Do(context.Background())
	var good model.Goods
	for _, hit := range result.Hits.Hits {
		good.UnmarshalJSON(*hit.Source)
	}
	return good
}

func GetGoodsByCategory(category, pageNum, pageSize int) []model.Goods {

	var goods []model.Goods
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("category_id", category)).Must(elastic.NewTermQuery("status", 1))
	result, err := es.Search().Index().Type().Size(pageSize).From(pageNum).Query(query).Do(context.Background())
	if err != nil {
		return goods
	}
	for _, hit := range result.Hits.Hits {
		good := model.Goods{}
		good.UnmarshalJSON(*hit.Source)
		goods = append(goods, good)
	}
	return goods
}
