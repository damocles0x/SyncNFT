package ES

import (
	"SyncNFT/config"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"time"
)

func SaveOrUpdateData(data *EsEnity) {
	result, err := GetByID4ES(context.Background(), "nft", data.ID)
	if result == "" && err != nil {
		data.CreatedTime = time.Now()
		data.UpdatedTime = time.Now()
		marshal, _ := json.Marshal(data)
		Create(context.Background(), "nft", data.ID, string(marshal))
	}
}

// GetByID4ES 根据ID查询单个文档
func GetByID4ES(ctx context.Context, index, id string) (string, error) {
	res, err := config.EsCli.Get().Index(index).Id(id).Do(ctx)
	if err != nil {
		return "", err
	}
	return string(res.Source), nil
}

func GetByAddress(ctx context.Context, index, address string) (*elastic.SearchResult, error) {
	termQuery := elastic.NewTermQuery("owner", address)
	searchResult, err := config.EsCli.Search().
		Index(index).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return searchResult, nil
}

func SaveOrUpsert(ctx context.Context, index, id string, doc interface{}) error {
	_, err := config.EsCli.Update().Index(index).Id(id).Doc(doc).Upsert(doc).Refresh("true").Do(ctx)
	return err
}
func Create(ctx context.Context, index, id, json string) error {
	_, err := config.EsCli.Index().Index(index).OpType("create").Id(id).BodyJson(json).Refresh("true").Do(ctx)
	return err
}
func Update(ctx context.Context, index, id string, doc interface{}) error {
	_, err := config.EsCli.Update().Index(index).Id(id).Doc(doc).Refresh("true").Do(ctx)
	return err
}

func ESIndexExists(ctx context.Context, index string) (bool, error) {
	return config.EsCli.IndexExists(index).Do(ctx)
}

func CrtESIndex(ctx context.Context, index string) error {
	mapping := `{
	"mappings":{
		"dynamic": "strict",
		"properties":{
			"id": 				{ "type": "keyword" },
			"token_id":			{ "type": "keyword" },
			"token_uri":				{ "type": "keyword" },
			"owner":		{ "type": "keyword" },
			"oracle_addr":         { "type": "keyword" },
			"token_approval": { "type": "text" },
			"updated_time":		{ "type": "date" },
			"created_time":		{ "type": "date" }
			}
		}
	}`
	exist, err := ESIndexExists(ctx, index)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	_, err = config.EsCli.CreateIndex(index).BodyString(mapping).Do(ctx)
	return err
}
