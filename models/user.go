package models

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	DOB          string `json:"dob"`
	FavoriteCity string `json:"favorite_city"`
	Admin        bool   `json:"admin"`
	AuthToken    string `json:"auth_token"`
	aaa          string `xorm:"aaa"`
}

func NewUSer() *User {
	return &User{}
}
