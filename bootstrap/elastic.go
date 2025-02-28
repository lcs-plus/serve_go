package bootstrap

import (
	"0121_1/global"
	"github.com/elastic/go-elasticsearch/v8"
)

func ElasticInit() *elasticsearch.Client {
	es, _ := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			global.App.Config.Elastic.Addr,
		},
		Username: global.App.Config.Elastic.Username,
		Password: global.App.Config.Elastic.Password,
	})
	return es
}
