package main

func SetIndex(shards, replicas string) string {
	return `{
    "settings": {
        "index": {
            "number_of_shards": ` + shards + `,
            "number_of_replicas": ` + replicas + `,
			"auto_expand_replicas": ` + replicas + `,
			"auto_expand_replicas": ` + replicas + `,
			"auto_expand_replicas": ` + replicas + `,
			"auto_expand_replicas": ` + replicas + `,
			"auto_expand_replicas": ` + replicas + `,
			"auto_expand_replicas": ` + replicas + `,

        }
    }
}`
}

const Order = `{
	"order":{
		"properties":{
			"order_id":{"type":"keyword"},
			"user_id":{"type":"keyword"},
			"goods":{"type":"nested","properties":{
				"goods_id":{"type":"integer"},
				"goods_num":{"type":"integer"},
				"goods_price":{"type":"double"},
				"goods_total_price":{"type":"double"},
			}},
		}
	}
}
`

const User = `{
	"user":{
		"properties":{
			"user_id":{"type":"keyword"},
"username":{"type":"keyword"},
""
		}
	}
}`
