package middleware

import (
	"WeatherServer/common"
	"WeatherServer/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.Request.Header.Get("Authorization")
		// token为空
		fmt.Println("token: ", tokenString)
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		// 非法token
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		// 提取token的有效部分:Bearer xxxxx
		tokenString = tokenString[7:]
		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		// 非法token
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		// 获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.Where("id =?", userId).First(&user)
		// 将用户信息写入上下文便于读取
		c.Set("user", user)
		c.Next()
	}
}
