package main

import (
	"final-assignment/config"
	"final-assignment/controller"
	"final-assignment/migrations"
	"final-assignment/repository"
	"final-assignment/routes"
	"final-assignment/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var (
		db         *gorm.DB           = config.SetUpDatabaseConnection()
		jwtService service.JWTService = service.NewJWTService()

		//user
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    service.UserService       = service.NewUserService(userRepository, jwtService)
		userController controller.UserController = controller.NewUserController(userService)

		//photo
		photoRepository repository.PhotoRepository = repository.NewPhotoRepository(db)
		photoService    service.PhotoService       = service.NewPhotoService(photoRepository, jwtService)
		photoController controller.PhotoController = controller.NewPhotoController(photoService)

		//comment
		commentRepository repository.CommentRepository = repository.NewCommentRepository(db)
		commentService    service.CommentService       = service.NewCommentService(commentRepository, jwtService)
		commentController controller.CommentController = controller.NewCommentController(commentService)

		//socialMedia
		socialMediaRepository repository.SocialMediaRepository = repository.NewSocialMediaRepository(db)
		socialMediaService    service.SocialMediaService       = service.NewSocialMediaService(socialMediaRepository, jwtService)
		socialMediaController controller.SocialMediaController = controller.NewSocialMediaController(socialMediaService)
	)

	server := gin.Default()
	routes.User(server, userController, jwtService)
	routes.Photo(server, photoController, jwtService)
	routes.Comment(server, commentController, jwtService)
	routes.SocialMedia(server, socialMediaController, jwtService)

	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("error migration: %v", err)
	}

	// not work properly
	// if err := migrations.Seeder(db); err != nil {
	// 	log.Fatalf("error migration seeder: %v", err)
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
