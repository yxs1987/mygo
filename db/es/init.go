package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mygo/common"
	"mygo/db"
	"mygo/model"
)

//是否添加示例数据
var example = flag.Bool("example", true, "-example=true")

//是否创建索引
var cindex = flag.Bool("cindex", true, "-cindex=true")

func main() {
	InitMapping(common.ES_GOODS, common.Goods)
	InitMapping(common.ES_ADDRESS, common.Address)
	InitMapping(common.ES_USER, common.User)
	InitMapping(common.ES_ORDER, common.Order)
}

// 创建 elasticSearch 的 Mapping
// es7已经去除type,本项目支持es6,默认index和type一致
func InitMapping(esIndexName string, typeMapping string) error {
	var err error
	indexMapping := SetIndex("1", "0")
	ctx := context.Background()
	client := db.ConnectEs()
	if err != nil {
		return err
	}
	_, err = client.DeleteIndex(esIndexName).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	// Create a new index.
	_, err = client.CreateIndex(esIndexName).BodyString(indexMapping).Do(ctx)
	if err != nil {
		log.Println("CreateIndex" + err.Error())
		return err
	}

	_, err = client.PutMapping().Index(esIndexName).Type(esIndexName).BodyString(typeMapping).Do(context.Background())

	if err != nil {
		log.Println("NewIndicesCreateService", err.Error(), esIndexName)
		return err
	}

	return err
}

//用例数据
func ExampleData(index string) {
	client := db.ConnectEs()

	switch index {
	case common.ES_GOODS:
		m := model.Goods{}
		m.GoodsId = common.GetNextId()
		m.GoodsName = "测试商品1"
		m.GoodsWeight = 0.3
		m.GoodsPrice = 125
		m.Image = []model.GoodsImage{}
		m.Content = "这里是商品的详细介绍"
		client.Index().Index(index).BodyJson(m).Do(context.Background())

	case common.ES_CATEGORY:
		c := model.Category{}
		c.Image = ""
		c.CategoryName = "分类1"
		c.ChildCategory = []model.Category{}
		client.Index().Index(index).BodyJson(c).Do(context.Background())

	}
}

func SetIndex(shards, replicas string) string {
	return `{
    	"settings": {
            "number_of_shards": ` + shards + `,
            "number_of_replicas": ` + replicas + `
		}
}`
}
