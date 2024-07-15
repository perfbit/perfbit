// pkg/model/user.go
package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Verified bool   `json:"verified"`
	Code     string `json:"code"`
}
