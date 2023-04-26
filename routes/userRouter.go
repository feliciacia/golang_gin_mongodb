package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", func(ctx *gin.Context) {})
	incomingRoutes.GET("/users/:users_id", func(ctx *gin.Context) {})
}
