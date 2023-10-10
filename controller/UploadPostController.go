package controller

import (
	"github.com/gin-gonic/gin"
	uuidGen "github.com/google/uuid"
	"os"
	"path/filepath"
	"sharephoto/utils/session"
)

func UploadPostController(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	User := session.SessionsManager.GetUser(ctx)
	if User == nil {
		ctx.Redirect(302, "/"+uuid)
		return
	}

	file, err := ctx.FormFile("img")
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	uuidImg := uuidGen.New().String()

	ext := filepath.Ext(file.Filename)

	dst, err := os.Create("src/cdn/" + uuidImg + ext)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
	}(dst)

	err = ctx.SaveUploadedFile(file, dst.Name())
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(302, "/"+uuid)

}
