package routes

import (
	"final-assignment/controller"
	"final-assignment/middleware"
	"final-assignment/service"

	"github.com/gin-gonic/gin"
)

func Comment(route *gin.Engine, commentController controller.CommentController, jwtService service.JWTService) {
	routes := route.Group("/comments", middleware.Authenticate(jwtService))
	{
		routes.POST("", commentController.PostComment)
		routes.GET("", commentController.GetAllComment)
		routes.GET("/:commentId", commentController.GetCommentById)
		routes.PUT("/:commentId", commentController.Update)
		routes.DELETE("/:commentId", commentController.Delete)
	}
}
