package dto

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}
