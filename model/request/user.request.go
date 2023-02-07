package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserUpdateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password,min=6,max=32"`
}
