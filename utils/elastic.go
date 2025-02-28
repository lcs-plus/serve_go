package utils

import (
	"0121_1/global"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func GetIndexList() (json.Decoder, error) {
	req := esapi.CatIndicesRequest{
		Format: "json",
	}

	res, err := req.Do(context.Background(), global.App.EsClient)

	if err != nil {
		return json.Decoder{}, err
	}

	defer res.Body.Close()

	return *json.NewDecoder(res.Body), nil

}
