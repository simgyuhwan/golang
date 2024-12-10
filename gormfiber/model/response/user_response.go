package response

type UserResponse struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}