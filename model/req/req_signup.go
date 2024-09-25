package req

type ReqSignUp struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"fullname" validate:"required"`
}
