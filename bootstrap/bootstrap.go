package bootstrap

import "0121_1/global"

func Init() {
	// 初始化配置
	InitializeConfig()
	// 初始化日志
	global.App.Log = InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = InitializeDB()

	//初始化redis
	global.App.Redis = InitRedis()

	//初始化es
	global.App.EsClient = ElasticInit()

	//初始化Casbin
	casbinInit()
}
