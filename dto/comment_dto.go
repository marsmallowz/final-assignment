package dto

type (
	CommentCreateRequest struct {
		Message string `json:"message" binding:"required"`
		PhotoID uint   `json:"photo_id" binding:"required"`
	}

	CommentResponse struct {
		ID      uint   `json:"id"`
		Message string `json:"message"`
		PhotoID uint   `json:"photo_id"`
		UserID  uint   `json:"user_id"`
	}

	GetAllCommentResponse struct {
		Comments []CommentResponse
	}

	CommentUpdateRequest struct {
		Message string `json:"message" binding:"required"`
	}
)
