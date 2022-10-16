package users

type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
