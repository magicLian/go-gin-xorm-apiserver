package controllers

import (
	"errors"
	"gin-test-go/dao"
	"gin-test-go/libs"
	"gin-test-go/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users, err := dao.GetUsers()
	if err != nil {
		log.Println("get users failed {}", err.Error())
		libs.HttpResult(c, http.StatusInternalServerError, err, nil)
	} else {
		libs.HttpResult(c, http.StatusOK, nil, users)
	}
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		libs.HttpResult(c, http.StatusBadRequest, err, nil)
		log.Println("parse user failed {}", err.Error())
	}
	if user.Email == "" {
		err := errors.New("email can not be empty")
		libs.HttpResult(c, http.StatusBadRequest, err, nil)
	}

	_, err = dao.CreateUser(user)

	if err != nil {
		libs.HttpResult(c, http.StatusExpectationFailed, err, nil)
	}

	libs.HttpResult(c, http.StatusOK, nil, user)
}
