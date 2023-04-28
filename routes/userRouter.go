package routes

import (
	controller "github.com/feliciacia/golang_gin_mongodb/golang_gin_mongodb/controllers"
	"github.com/feliciacia/golang_gin_mongodb/golang_gin_mongodb/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:users_id", controller.GetUser())
}
