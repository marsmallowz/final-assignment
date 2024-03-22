package dto

type (
	// kalau gak pakai binding bisa gak diisi saat request
	PhotoCreateRequest struct {
		Title    string `json:"title" binding:"required"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url" binding:"url"`
	}

	PhotoResponse struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url"`
		UserID   uint   `json:"user_id"`
	}

	UserAtPhotoResponse struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	GetAllPhotoResponse struct {
		Photos []PhotoResponse
	}

	PhotoUpdateRequest struct {
		Title    string `json:"title" binding:"required"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url" binding:"url"`
	}
)
