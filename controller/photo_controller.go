package controller

import (
	"final-assignment/dto"
	"final-assignment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	PhotoController interface {
		PostPhoto(ctx *gin.Context)
		GetAllPhoto(ctx *gin.Context)
		GetPhotoById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	photoController struct {
		photoService service.PhotoService
	}
)

func NewPhotoController(us service.PhotoService) PhotoController {
	return &photoController{
		photoService: us,
	}
}

func (c *photoController) PostPhoto(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uint)

	var photo dto.PhotoCreateRequest
	if err := ctx.ShouldBind(&photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.photoService.PostPhoto(ctx.Request.Context(), photo, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *photoController) GetAllPhoto(ctx *gin.Context) {
	result, err := c.photoService.GetAllPhoto(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result.Photos)
}

func (c *photoController) GetPhotoById(ctx *gin.Context) {
	photoId := ctx.Params.ByName("photoId")

	result, err := c.photoService.GetPhotoById(ctx.Request.Context(), photoId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *photoController) Update(ctx *gin.Context) {
	photoId := ctx.Params.ByName("photoId")

	var req dto.PhotoUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.MustGet("user_id").(uint)
	result, err := c.photoService.UpdatePhoto(ctx.Request.Context(), req, photoId, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *photoController) Delete(ctx *gin.Context) {
	photoId := ctx.Params.ByName("photoId")
	userId := ctx.MustGet("user_id").(uint)
	if err := c.photoService.DeletePhoto(ctx.Request.Context(), photoId, userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
