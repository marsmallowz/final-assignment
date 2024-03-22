package dto

type (
	SocialMediaCreateRequest struct {
		Name           string `json:"name" binding:"required"`
		SocialMediaURL string `json:"social_media_url" binding:"url"`
	}

	SocialMediaResponse struct {
		ID             uint   `json:"id"`
		Name           string `json:"name"`
		SocialMediaURL string `json:"social_media_url"`
		UserID         uint   `json:"user_id"`
	}

	GetAllSocialMediaResponse struct {
		SocialMedias []SocialMediaResponse
	}

	SocialMediaUpdateRequest struct {
		Name           string `json:"name" binding:"required"`
		SocialMediaURL string `json:"social_media_url" binding:"url"`
	}
)
