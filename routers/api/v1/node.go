package v1

import (
	"github.com/gin-gonic/gin"
	"gptc-sync/pkg/e"
	"net/http"
)

func CurrentNodeBlockNumber(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "1",
	})
}

func TotalNodeList(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "1",
	})
}

func HashRate(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "1",
	})
}
