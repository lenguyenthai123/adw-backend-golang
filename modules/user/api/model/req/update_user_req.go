package req

type UpdateUserReq struct {
	FullName string `json:"full_name" extensions:"x-order=1"`
	Email    string `json:"email" extensions:"x-order=2"`
	Phone    string `json:"phone" extensions:"x-order=3"`
}
