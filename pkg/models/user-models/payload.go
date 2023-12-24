package user_models

import (
	"github.com/dgrijalva/jwt-go"
)

type Payload struct {
	UserId int    `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
