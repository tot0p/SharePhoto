package controller

import "github.com/gin-gonic/gin"

func UploadApiController(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "pong",
	})

}
