package sales

import (
	"0121_1/app/controller"
	"0121_1/app/models/v1/sales"
	v1 "0121_1/app/services/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodayInfo(c *gin.Context) {
	var listForm sales.GetTodayListForm
	err := c.ShouldBind(&listForm)
	if err != nil {
		panic(err)
	}
	list := v1.GetList(&listForm)

	c.JSON(http.StatusOK, controller.Response{
		Code: 200,
		Data: list,
		Msg:  "ok",
	})
}
