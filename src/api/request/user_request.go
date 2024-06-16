package request

type UserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
