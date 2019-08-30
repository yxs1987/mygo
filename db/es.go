package db

import (
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
	"sync"
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
