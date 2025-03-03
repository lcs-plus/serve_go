package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {

		user := c.GetHeader("X-User")
		obj := c.Request.URL.Path
		act := c.Request.Method

		enforce, err := enforcer.Enforce(user, obj, act)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !enforce {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}
		// 权限检查通过，继续处理请求
		c.Next()
	}
}
