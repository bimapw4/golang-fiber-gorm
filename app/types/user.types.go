package types

type UserDTO struct {
	Username string `json:"username"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type UserUpdateDTO struct {
	Email string `json:"email"`
}
