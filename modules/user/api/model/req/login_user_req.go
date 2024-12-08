package req

type LoginUserReq struct {
	Email    string `json:"email" validate:"required,email" extensions:"x-order=1"`
	Password string `json:"password" validate:"required" extensions:"x-order=2"`
}
