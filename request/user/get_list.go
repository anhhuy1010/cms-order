package user

type (
	GetListRequest struct {
		Keyword  string  `form:"keyword"`
		Username *string `form:"username"`
		Page     int     `form:"page"`
		Limit    int     `form:"limit"`
		Sort     string  `form:"sort"`
		IsActive *int    `form:"is_active" `
	}
	ListResponse struct {
		Uuid       string `json:"uuid" `
		ClientUuid string `json:"client_uuid"`
		Name       string `json:"name"`
		UserName   string `json:"username"`
		IsActive   int    `json:"is_active"`
	}
)
