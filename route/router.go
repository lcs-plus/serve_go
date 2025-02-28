package route

import (
	"0121_1/app/controller"
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

	r.GET("/get-list", controller.GetList)

	r.GET("/get-es-list", controller.GetTestList)

	return r
}
