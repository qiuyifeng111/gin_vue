package middleware

import (
	"fmt"
	"goforpra/common"
	"goforpra/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		fmt.Println(tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		fmt.Println(token)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
