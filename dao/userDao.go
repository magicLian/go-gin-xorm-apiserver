package dao

import (
	"errors"
	"gin-test-go/db"
	"gin-test-go/models"
)

func GetUsers() ([]models.User, error) {
	userList := []models.User{}
	err := db.GetDbInstance().Find(&userList)
	if err != nil {
		return nil, err
	}
	return userList, err
}

func CreateUser(u *models.User) (int64, error) {
	i, err := db.GetDbInstance().Insert(u)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func UpdateUser(u models.User) (int64, error) {
	i, err := db.GetDbInstance().Where("id=", u.Id).Update(u)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func DeleteUserById(id int64) (int64, error) {
	b, err := IsUserExistsById(id)
	if err != nil {
		return 0, err
	}
	if !b {
		return 0, errors.New("user is not exist")
	}

	user := &models.User{Id: id}
	i, err := db.GetDbInstance().Delete(user)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func IsUserExistsById(id int64) (bool, error) {
	user := &models.User{Id: id}
	has, err := db.GetDbInstance().Exist(user)
	return has,err
}
