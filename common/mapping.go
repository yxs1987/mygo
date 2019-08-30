package common

const Order = `{
	"es_order":{
		"properties":{
			"order_id":{"type":"keyword"},
			"user_id":{"type":"keyword"},
			"goods":{
				"type":"nested",
				"properties":{
					"goods_id":{"type":"keyword"},
					"goods_num":{"type":"integer"},
					"goods_price":{"type":"double"},
					"goods_total_price":{"type":"double"},
					"goods_weight":{"type":"double"}
				}
			},
			"total_num":{"type":"double"},
			"total_weight":{"type":"double"},
			"total_price":{"type":"double"}
		}
	}
}
`

const User = `{
	"es_user":{
		"properties":{
			"user_id":{"type":"keyword"},
			"nickname":{"type":"text"},
			"name":{"type":"text"},
			"open_id":{"type":"keyword"},
			"address":{
				"type":"nested",
				"properties":{
					"address_id":{"type":"keyword"},
					"province":{"type":"keyword"},
					"city":{"type":"keyword"},
					"district":{"type":"keyword"},
					"detail":{"type":"text"},
					"mobile":{"type":"integer"},
					"name":{"type":"keyword"}
				}
			},
			"avatarUrl":{"type":"keyword"},
			"city":{"type":"keyword"},
			"country":{"type":"keyword"},
			"province":{"type":"keyword"},
			"language":{"type":"keyword"},
			"gender":{"type":"integer"}
		}
	}
}`

const Address = `{
	"es_address":{
		"properties":{
			"user_id":{"type":"keyword"},
			"address_id":{"type":"text"},
			"status":{"type":"integer"},
			"province":{"type":"keyword"},
			"city":{"type":"keyword"},
			"district":{"type":"keyword"},
			"detail":{"type":"text"},
			"mobile":{"type":"long"},
			"name":{"type":"keyword"}
		}
	}
}`

const Goods = `{
	"es_goods":{
		"properties":{
			"order_id":{"type":"keyword"},
			"user_id":{"type":"keyword"},
			"goods":{
				"type":"nested",
				"properties":{
					"goods_id":{"type":"keyword"},
					"goods_num":{"type":"integer"},
					"goods_price":{"type":"double"},
					"goods_total_price":{"type":"double"},
					"goods_weight":{"type":"double"}
				}
			},
			"total_num":{"type":"double"},
			"total_weight":{"type":"double"},
			"total_price":{"type":"double"}
		}
	}
}`
