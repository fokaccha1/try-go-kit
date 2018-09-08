package user

type UserDraft struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GetUserRequest struct {
	Id int `json:"id"`
}

type GetUserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUserResponse struct {
	Id int `json:"id"`
}
