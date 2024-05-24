package user

type (
	UpdateUri struct {
		Uuid string `uri:"uuid"`
	}
	UpdateRequest struct {
		Name     string `json:"name"`
		UserName string `json:"username"`
		Email    string `json:"email"`
	}
)
