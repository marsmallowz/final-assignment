package routes

import (
	"final-assignment/controller"
	"final-assignment/middleware"
	"final-assignment/service"

	"github.com/gin-gonic/gin"
)

func SocialMedia(route *gin.Engine, socialMediaController controller.SocialMediaController, jwtService service.JWTService) {
	routes := route.Group("/socialmedias", middleware.Authenticate(jwtService))
	{
		routes.POST("", socialMediaController.CreateSocialMedia)
		routes.GET("", socialMediaController.GetAllSocialMedia)
		routes.GET("/:socialMediaId", socialMediaController.GetSocialMediaById)
		routes.PUT("/:socialMediaId", socialMediaController.Update)
		routes.DELETE("/:socialMediaId", socialMediaController.Delete)
	}
}
