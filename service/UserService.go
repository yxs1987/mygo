package service

import (
	"context"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/olivere/elastic"
	"github.com/uniplaces/carbon"
	"mygo/common"
	"mygo/db"
	"mygo/model"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

var es *elastic.Client

func init() {
	es = db.ConnectEs()
}

func CreateUser(openid string, info model.UserInfo) Response {
	var res Response
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("open_id", openid))

	result, err := es.Search().Index(common.ES_USER).Size(1).Query(query).Do(context.Background())
	if err != nil {
		res.StatusCode = 400
		res.Msg = "查询错误"
		return res
	}

	var user model.UserInfo
	res.StatusCode = 200
	res.Data = user

	for _, hit := range result.Hits.Hits {
		user.UnmarshalJSON(*hit.Source)
	}
	if user.OpenId == "" {
		currWorker := &idworker.IdWorker{}
		currWorker.InitIdWorker(1000, 1)
		newId, err := currWorker.NextId()
		if err != nil {
			panic(err)
		}

		user = info
		user.UserId = newId
		user.LatestLoginTime = carbon.Now().Format("2006-01-02 15:04:05")
		_, err = es.Index().Index(common.ES_USER).Type(common.ES_USER).BodyJson(user).Do(context.Background())
		if err != nil {
			res.StatusCode = 400
			res.Msg = "用户数据新增失败"
		}
	} else {
		up := map[string]interface{}{}
		up["latest_login_time"] = carbon.Now().Format("2006-01-02 15:04:05")
		up["avatar_url"] = info.AvatarUrl
		up["nickName"] = info.Nickname
		up["country"] = info.Country
		up["province"] = info.Province
		up["city"] = info.City
		_, err = es.Update().Index(common.ES_USER).Type(common.ES_USER).Doc(up).Do(context.Background())
		if err != nil {
			res.StatusCode = 400
			res.Msg = "用户数据更新失败"

		}
	}

	return res
}

func GetUserById(user_id int64) Response {
	query := elastic.NewTermQuery("user_id", user_id)

	var user model.UserInfo
	result, _ := es.Search().Index(common.ES_USER).Type(common.ES_USER).Query(query).Size(1).Do(context.Background())
	for _, hit := range result.Hits.Hits {
		user.UnmarshalJSON(*hit.Source)
	}

	res := Response{
		StatusCode: 200,
		Msg:        "",
		Data:       user,
	}
	return res
}
