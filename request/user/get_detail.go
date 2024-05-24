package user

type (
	GetDetailUri struct {
		Uuid string `uri:"uuid"`
	}
	GetDetailResponse struct {
		Uuid     string `json:"uuid" `
		Name     string `json:"name"`
		UserName string `json:"username"`
		Email    string `json:"email"`
	}
)
