package req

type ChangePasswordReq struct {
	OldPwd string `json:"old_pwd" validate:"required" extensions:"x-order=1"`
	NewPwd string `json:"new_pwd" validate:"required" extensions:"x-order=2"`
}
