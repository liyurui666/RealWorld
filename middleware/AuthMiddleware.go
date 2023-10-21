package middleware

import (
	"RealWorld/common"
	"RealWorld/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无该权限",
			})
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		//jwt解析失败
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无该权限",
			})
			return
		}
		email := claims.Email
		var user models.User
		user.Email = email
		u := models.QueryEmail(user)
		//用户不存在
		if u.Id == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无该权限",
			})
			return
		}

		context.Set("user", u)
		context.Next()
	}
}
