package types

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
}
