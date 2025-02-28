package global

import (
	"0121_1/config"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config   config.Configuration
	DB       *gorm.DB
	Log      *zap.Logger
	Redis    *redis.Client
	EsClient *elasticsearch.Client
}

var App = new(Application)
