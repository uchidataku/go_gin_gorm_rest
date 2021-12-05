package server

import (
	"github.com/gin-gonic/gin"
	"go_gin_gorm_rest/controller"
)

func Init() {
	r := router()
	r.Run(":3000")
}

func router() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(cc *gin.Context) {
		cc.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.Login)
	}

	a := r.Group("/accounts")
	{
		ctrl := controller.Controller{}
		a.GET("", ctrl.Index)
		a.GET("/:id", ctrl.Show)
		a.POST("", ctrl.Create)
		a.PUT("/:id", ctrl.Update)
		a.DELETE("/:id", ctrl.Delete)
	}

	return r
}
