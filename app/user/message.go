package user

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GetUserRequest struct {
	Id int `json:"id"`
}
