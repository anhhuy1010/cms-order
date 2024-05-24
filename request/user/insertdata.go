package user

type (
	GetInsertRequest struct {
		Uuid       string `json:"uuid" `
		ClientUuid string `json:"client_uuid"`
		Name       string `json:"name"`
		UserName   string `json:"username"`
		IsActive   int    `json:"is_active"`
	}
	InsertResponse struct {
		Uuid       string `json:"uuid" `
		ClientUuid string `json:"client_uuid"`
		Name       string `json:"name"`
		UserName   string `json:"username"`
		IsActive   int    `json:"is_active"`
	}
)
