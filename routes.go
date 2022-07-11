package main

import (
	"goforpra/controller"
	"goforpra/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("register", controller.Register)
	r.POST("login", controller.Login)
	r.POST("info", middleware.AuthMiddleware(), controller.Info)
	return r
}
