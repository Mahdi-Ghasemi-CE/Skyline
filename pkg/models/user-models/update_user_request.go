package user_models

type UpdateUserRequest struct {
	UserId     int    `json:"userId" validate:"required"`
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"rePassword" validate:"required"`
}
