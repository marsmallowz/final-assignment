package dto

type (
	UserCreateRequest struct {
		Email           string `json:"email" binding:"email"`
		Username        string `json:"username" binding:"required"`
		Password        string `json:"password" binding:"required,min=6"`
		Age             uint8  `json:"age" binding:"required,gte=8"`
		ProfileImageURL string `json:"profile_image_url" binding:"omitempty,url"`
	}

	UserResponse struct {
		ID              uint   `json:"id"`
		Email           string `json:"email"`
		Username        string `json:"username"`
		Age             uint8  `json:"age"`
		ProfileImageURL string `json:"profile_image_url"`
	}

	UserUpdateRequest struct {
		Email           string `json:"email" binding:"email"`
		Username        string `json:"username" binding:"required"`
		Age             uint8  `json:"age" binding:"required,gte=8"`
		ProfileImageURL string `json:"profile_image_url" binding:"omitempty,url"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" binding:"email"`
		Password string `json:"password" binding:"required"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
	}
)
