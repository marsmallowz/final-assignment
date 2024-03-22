package routes

import (
	"final-assignment/controller"
	"final-assignment/middleware"
	"final-assignment/service"

	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/users")
	{
		routes.POST("/register", userController.Register)
		routes.POST("/login", userController.Login)
		routes.PUT("/", middleware.Authenticate(jwtService), userController.Update)
		routes.DELETE("/", middleware.Authenticate(jwtService), userController.Delete)
	}
}
