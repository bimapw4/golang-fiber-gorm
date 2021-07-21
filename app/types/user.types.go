package types

type UserDTO struct {
	Username string `json:"username"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type UserUpdateDTO struct {
	Email string `json:"email"`
}

type UserListDB struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type UserListResp struct {
	Status bool         `json:"status"`
	Data   []UserListDB `json:"data"`
}
