package route

import (
	"0121_1/app/controller"
	"0121_1/app/controller/v1/sales"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	//r.Use(middleware.CasbinMiddleware(global.App.Casbin))
	r.POST("/api/get-list", controller.GetList)

	r1 := r.Group("/api")
	{
		saleRoute := r1.Group("/sales")
		{
			saleRoute.POST("/get-today-info", sales.GetTodayInfo)
		}
	}
	return r
}
