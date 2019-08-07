package controllers

import (
	"gin-test-go/dao"
	"gin-test-go/libs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users, err := dao.GetUsers()
	if err != nil {
		log.Println("get users failed {}",err.Error())
		libs.HttpResult(c,http.StatusInternalServerError,err,nil)
	} else {
		libs.HttpResult(c,http.StatusOK,nil,users)
	}
}

func CreateUser(c *gin.Context)  {

}
