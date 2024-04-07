package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
