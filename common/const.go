package common

import idworker "github.com/gitstliu/go-id-worker"

const ES_USER = "es_user"
const ES_ORDER = "es_order"
const ES_ADDRESS = "es_address"
const ES_GOODS = "es_goods"
const ES_CATEGORY = "es_category"
const ES_CART = "es_cart"

func GetNextId() int64 {
	currWorker := &idworker.IdWorker{}
	currWorker.InitIdWorker(1000, 1)
	newId, _ := currWorker.NextId()
	return newId
}
