package models

import "github.com/dgrijalva/jwt-go"

type AccountInfo struct {
	ID        uint   `json:"-" gorm:"primaryKey"`
	Username  string `json:"username" `
	Password  string `json:"password"`
	Role      string `json:"role"`
	BaseModel `json:"-"`
}

type AccountClaims struct {
	jwt.StandardClaims
	Name     string
	Identity string
}
