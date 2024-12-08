package req

type ResetPasswordReq struct {
	Email string `json:"email" validate:"required,email"`
}
