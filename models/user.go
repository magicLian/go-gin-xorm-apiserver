package models

type User struct {
	ID           int64  `xorm:"id"`
	Email        string `xorm:"email"`
	DOB          string `xorm:"dob"`
	FavoriteCity string `xorm:"favorite_city"`
	Admin        bool   `xorm:"admin"`
	AuthToken    string `xorm:"auth_token"`
	aaa          string `xorm:"aaa"`
}

func NewUSer() *User {
	return &User{}
}
