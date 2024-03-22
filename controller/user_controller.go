package controller

import (
	"final-assignment/dto"
	"final-assignment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user dto.UserCreateRequest
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.userService.RegisterUser(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *userController) Login(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.userService.Verify(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *userController) Update(ctx *gin.Context) {
	var req dto.UserUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.MustGet("user_id").(uint)

	result, err := c.userService.UpdateUser(ctx.Request.Context(), req, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *userController) Delete(ctx *gin.Context) {
	userIdFromToken := ctx.MustGet("user_id").(uint)

	if err := c.userService.DeleteUser(ctx.Request.Context(), userIdFromToken); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
