package db

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"

	"os"
	"sync"
	"ytc_middle_cloud/module/model/es"
)

type Index struct {
	NumberOfShards       int  `json:"number_of_shards"`
	RoutingPartitionSize int  `json:"routing_partition_size"`
	NumberOfReplicas     int  `json:"number_of_replicas"`
	AutoExpandReplicas   bool `json:"auto_expand_replicas"`
	RefreshInterval      int  `json:"refresh_interval"`
	MaxResultWindow      int  `json:"max_result_window"`
}

var once sync.Once
var client *elastic.Client

func ConnectEs() *elastic.Client {
	var connString string

	once.Do(func() {

		connString = fmt.Sprintf("http://%s:%d", "192.168.99.100", 9200)

		esClient, err := elastic.NewClient(
			elastic.SetSniff(false),
			elastic.SetURL(connString),
			elastic.SetBasicAuth("", ""),
			elastic.SetGzip(true),
			elastic.SetErrorLog(log.New(os.Stderr, "es-err", log.LstdFlags)),
			elastic.SetInfoLog(log.New(os.Stdout, "es-info", log.LstdFlags)),
			elastic.SetTraceLog(log.New(os.Stdin, "es-trace", log.LstdFlags)))

		if err != nil {
			panic(err)
		}
		client = esClient
	})

	return client
}

// 创建 elasticSearch 的 Mapping
// es7已经去除type,本项目支持es6,默认index和type一致
func InitMapping(esIndexName string, typeMapping string) error {
	var err error
	indexMapping := es.GetShard("1")
	ctx := context.Background()
	client := ConnectEs()
	if err != nil {
		return err
	}
	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(esIndexName).Do(ctx)
	if err != nil {
		log.Println("IndexExists" + err.Error())
		return err
	}
	//log.Println("es index: " + esIndexName)
	//log.Println("es type: " + esTypeName)
	//log.Println("es index mapping: " + indexMapping)
	//log.Println("es type mapping: " + typeMapping)
	if !exists {
		log.Println("es index not exists: " + esIndexName)
		// Create a new index.
		createIndex, err := client.CreateIndex(esIndexName).Body(indexMapping).Do(ctx)
		if err != nil {
			log.Println("CreateIndex" + err.Error())
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			//return errors.New("create index:" + esIndexName + ", not Ack nowledged")
		}
	}
	/**
	  * 判断 type 是否存在
	     exists, err = client.TypeExists().Index(esIndexName).Type(esTypeName).Do(ctx)
	     if err != nil {
	         return err
	     }
	     if !exists {

	     }
	*/
	// PutMapping() *IndicesPutMappingService

	putresp, err := client.PutMapping().Index(esIndexName).Type(esIndexName).BodyString(typeMapping).Do(context.Background())
	// 新建 mapping
	//indicesCreateResult, err := elastic.NewIndicesCreateService(client).Index(esIndexName).BodyString(mapping).Do(ctx)
	if err != nil {
		log.Println("NewIndicesCreateService" + err.Error())
		return err
	}
	if !putresp.Acknowledged {
		// Not acknowledged
		//return errors.New("create mapping fail, esIndexName:" + esIndexName + ", not Ack nowledged")
	}

	// 插入数据
	/*
		    type WholeBrowserData struct {
		        BrowserId     string                `json:"browser_id"`
		        BrowserName  string                `json:"browser_name"`
		    }

		    // Index a tweet (using JSON serialization)
			wholeBrowserData := WholeBrowserData{BrowserId: "BrowserId", BrowserName: "BrowserName" }
			put1, err := client.Index().
				Index(esIndexName).
				Type(esTypeName).
				Id("1").
				BodyJson(wholeBrowserData).
				Do(ctx)
			if err != nil {
				// Handle error
				panic(err)
			}
			log.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	*/

	return err
}

//用例数据
func ExampleData() {

}
