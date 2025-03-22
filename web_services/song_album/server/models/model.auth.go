// go get github.com/golang-jwt/jwt/v5
package models

import "github.com/golang-jwt/jwt/v5"

// Claims structure
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type Auth struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	IsDeleted bool   `json:"isdeleted"` // to soft delete the user
	Project   string `json:"project"`   // to identify the project for which user is created as this is shared table
}
