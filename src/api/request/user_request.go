package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email,min=4,max=200"`
	Name     string `json:"name" binding:"required,min=4,max=200"`
	Password string `json:"password" binding:"required,min=8,max=200"`
}
