package req

type CreateUserReq struct {
	FullName string `json:"full_name" validate:"required" extensions:"x-order=1"`
	Email    string `json:"email" validate:"required,email" extensions:"x-order=2"`
	Phone    string `json:"phone" validate:"phone" extensions:"x-order=3"`
	Password string `json:"password" validate:"required,password" extensions:"x-order=4"`
}
