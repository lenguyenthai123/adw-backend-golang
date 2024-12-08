package constant

type JWTPayload struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
