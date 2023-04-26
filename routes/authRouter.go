package routes

import (
	"github.com/feliciacia/golang_gin_mongodb/golang_gin_mongodb/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUp())
	incomingRoutes.POST("users/login", func(ctx *gin.Context) {})
}
