package dto

type User struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Age        int    `json:"age"`
	Permission string `json:"permission"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type UpdateUserReq struct {
	Username   string `json:"username"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Password   string `json:"password"`
	Permission string `json:"permission"`
}

type ListUserResponse struct {
	Data []*User `json:"data"`
}

type ListUserReq struct {
	Limit   string `json:"limit"`
	Offset  string `json:"offset"`
	OrderBy string `json:"orderBy"`
}
