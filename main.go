package main

import (
	"0121_1/bootstrap"
	"0121_1/global"
	"0121_1/route"
)

func main() {

	//初始化
	bootstrap.Init()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	r := route.InitRouter()

	r.Run(":" + global.App.Config.HttpApp.Port)
}
