package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username" `
	Password string `json:"password" `
}

type MyClaims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RespToken struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   Token  `json:"data"`
}

type RespLogin struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   Login  `json:"data"`
}

type Login struct {
	Uid          int    `json:"uid"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
