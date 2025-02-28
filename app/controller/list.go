package controller

import (
	"0121_1/app/models"
	"0121_1/app/services"
	"0121_1/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetList(c *gin.Context) {
	var listForm models.GetListForm
	err := c.ShouldBind(&listForm)
	if err != nil {
		return
	}

	list := services.GetList(&listForm)

	c.JSON(http.StatusOK, Response{
		200,
		list,
		"ok",
	})
}

func AddMsg(c *gin.Context) {
	var addMsg models.AddListForm
	err := c.ShouldBind(&addMsg)
	if err != nil {
		return
	}

}

func GetTestList(c *gin.Context) {
	list, _ := utils.GetIndexList()

	c.JSON(http.StatusOK, Response{
		200,
		list,
		"ok",
	})
}
