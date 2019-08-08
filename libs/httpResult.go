package libs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpResult(c *gin.Context,code int, err error, data interface{})  {
	if code == 200{
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":  data,
		})
	}

	if code == 400{
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	}

	if code == 404 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
			"data":    "",
		})
	}

	if code == 405 {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"status":  http.StatusMethodNotAllowed,
			"message": "Method not allow",
			"data":    "",
		})
	}

	if code == 403 {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": err.Error(),
			"data":    "",
		})
	}

	if code == 500 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
	}
}
