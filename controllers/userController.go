package controllers

import (
	"gin-test-go/dao"
	"gin-test-go/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUser(c *gin.Context) {
	m := new(models.User)

	if m1, err := dao.GetUser(m); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    m1,
		})
	}
}
