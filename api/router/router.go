package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reaovyd/learngorestapi/api/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.DisplayRoot)
	r.POST("/", controllers.ProcessRoot)
}
