package routes

import (
	"final-assignment/controller"
	"final-assignment/middleware"
	"final-assignment/service"

	"github.com/gin-gonic/gin"
)

func Photo(route *gin.Engine, photoController controller.PhotoController, jwtService service.JWTService) {
	routes := route.Group("/photos", middleware.Authenticate(jwtService))
	{
		routes.POST("", photoController.PostPhoto)
		routes.GET("", photoController.GetAllPhoto)
		routes.GET("/:photoId", photoController.GetPhotoById)
		routes.PUT("/:photoId", photoController.Update)
		routes.DELETE("/:photoId", photoController.Delete)
	}
}
