package controller

import "github.com/gin-gonic/gin"

func IndexController(ctx *gin.Context) {

	//uuid := ctx.Param("uuid")

	ctx.HTML(200, "index.html", gin.H{})
}
