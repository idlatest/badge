package users

import (
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func Init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func GetTokenAuth() *jwtauth.JWTAuth {
	return tokenAuth
}
