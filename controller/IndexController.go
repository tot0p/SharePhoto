package controller

import (
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
)

func IndexController(ctx *gin.Context) {

	uuid := uuid2.New().String()

	ctx.Redirect(302, "/"+uuid)
}
