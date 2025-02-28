package controller

import (
	"0121_1/app/models"
	"0121_1/app/services"
	"0121_1/global"
	"0121_1/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Register(c *gin.Context) {

	var user models.UserForm

	err := c.ShouldBind(&user)
	if err != nil {
		panic(err)
	}

	services.Register(&user)

	c.JSON(200, gin.H{
		"name": user.Name,
		"pass": user.Password,
	})
}

func Login(c *gin.Context) {
	var user models.UserForm
	err := c.ShouldBind(&user)
	if err != nil {
		panic(err)
	}

	result := services.Login(&user)

	if result != nil {

		accessToken, err := utils.GetToken(result.ID.ID, result.Name)
		if err != nil {
			panic(err)
		}

		ctx := context.Background()

		key := strconv.Itoa(result.ID.ID)

		global.App.Redis.Set(ctx, key, accessToken, time.Duration(global.App.Config.Jwt.ExpiresAt)*time.Second)

		c.JSON(http.StatusOK, Response{
			2000,
			models.UserLoginResult{User: *result, AccessToken: accessToken},
			"ok",
		})

	} else {
		c.JSON(400, gin.H{
			"msg": "用户名或密码错误",
		})
	}
}
