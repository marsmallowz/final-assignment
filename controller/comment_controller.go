package controller

import (
	"final-assignment/dto"
	"final-assignment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	CommentController interface {
		PostComment(ctx *gin.Context)
		GetAllComment(ctx *gin.Context)
		GetCommentById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	commentController struct {
		commentService service.CommentService
	}
)

func NewCommentController(us service.CommentService) CommentController {
	return &commentController{
		commentService: us,
	}
}

func (c *commentController) PostComment(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uint)

	var comment dto.CommentCreateRequest
	if err := ctx.ShouldBind(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.commentService.PostComment(ctx.Request.Context(), comment, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *commentController) GetAllComment(ctx *gin.Context) {
	result, err := c.commentService.GetAllComment(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result.Comments)
}

func (c *commentController) GetCommentById(ctx *gin.Context) {
	commentId := ctx.Params.ByName("commentId")

	result, err := c.commentService.GetCommentById(ctx.Request.Context(), commentId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *commentController) Update(ctx *gin.Context) {
	commentId := ctx.Params.ByName("commentId")

	var req dto.CommentUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.MustGet("user_id").(uint)
	result, err := c.commentService.UpdateComment(ctx.Request.Context(), req, commentId, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *commentController) Delete(ctx *gin.Context) {
	commentId := ctx.Params.ByName("commentId")
	userId := ctx.MustGet("user_id").(uint)
	if err := c.commentService.DeleteComment(ctx.Request.Context(), commentId, userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
