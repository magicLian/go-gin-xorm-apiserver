package models

import "time"

type User struct {
	Id           int64
	Email        string `xorm:"email"`
	DOB          string `xorm:"dob"`
	FavoriteCity string `xorm:"favorite_city"`
	Admin        bool   `xorm:"admin"`
	AuthToken    string `xorm:"auth_token"`
	Aaa          string `xorm:"aaa"`
	CreateTime   time.Time `xorm:"created"`
	UpdateTime   time.Time `xorm:"updated"`
}

func NewUSer() *User {
	return &User{}
}
