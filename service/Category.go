package service

import (
	"context"
	"github.com/olivere/elastic"
	"mygo/common"
	"mygo/model"
)

func GetCategory(id int64) Response {
	query := elastic.NewBoolQuery().MustNot(elastic.NewTermQuery("status", 13)).Must(elastic.NewTermQuery("id", id))
	result, err := es.Search().Index(common.ES_CATEGORY).Size(1).Query(query).Do(context.Background())
	if err != nil {

	}
	var cat model.Category
	for _, hit := range result.Hits.Hits {
		cat.UnmarshalJSON(*hit.Source)
		resp.Data = cat
	}
	return resp
}

func InsertCategory(data model.Category) error {
	_, err := es.Index().Index(common.ES_CATEGORY).BodyJson(data).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
