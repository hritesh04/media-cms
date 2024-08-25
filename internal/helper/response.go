package helper

import "github.com/gin-gonic/gin"

func ReturnSuccess(g *gin.Context, status int, data interface{}) {
	g.JSON(status, gin.H{
		"success": true,
		"data":    data,
	})
}

func ReturnFailed(g *gin.Context, status int, data interface{}) {
	g.JSON(status, gin.H{
		"success": false,
		"data":    data,
	})
}
