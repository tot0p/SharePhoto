package controller

import "github.com/gin-gonic/gin"

func GetApiController(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "pong",
	})

}
