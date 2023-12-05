package user_models

type UserResponse struct {
	UserId    int    `json:"-"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	RoleId    int    `json:"-"`
}
