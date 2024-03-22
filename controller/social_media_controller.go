package controller

import (
	"final-assignment/dto"
	"final-assignment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	SocialMediaController interface {
		CreateSocialMedia(ctx *gin.Context)
		GetAllSocialMedia(ctx *gin.Context)
		GetSocialMediaById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	socialMediaController struct {
		socialMediaService service.SocialMediaService
	}
)

func NewSocialMediaController(us service.SocialMediaService) SocialMediaController {
	return &socialMediaController{
		socialMediaService: us,
	}
}

func (c *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uint)

	var socialMedia dto.SocialMediaCreateRequest
	if err := ctx.ShouldBind(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.socialMediaService.CreateSocialMedia(ctx.Request.Context(), socialMedia, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *socialMediaController) GetAllSocialMedia(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uint)

	result, err := c.socialMediaService.GetAllSocialMedia(ctx.Request.Context(), userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result.SocialMedias)
}

func (c *socialMediaController) GetSocialMediaById(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uint)
	socialMediaId := ctx.Params.ByName("socialMediaId")

	result, err := c.socialMediaService.GetSocialMediaById(ctx.Request.Context(), socialMediaId, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *socialMediaController) Update(ctx *gin.Context) {
	socialMediaId := ctx.Params.ByName("socialMediaId")

	var req dto.SocialMediaUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.MustGet("user_id").(uint)
	result, err := c.socialMediaService.UpdateSocialMedia(ctx.Request.Context(), req, socialMediaId, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *socialMediaController) Delete(ctx *gin.Context) {
	socialMediaId := ctx.Params.ByName("socialMediaId")
	userId := ctx.MustGet("user_id").(uint)
	if err := c.socialMediaService.DeleteSocialMedia(ctx.Request.Context(), socialMediaId, userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your socialMedia has been successfully deleted",
	})
}
