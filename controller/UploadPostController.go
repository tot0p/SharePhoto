package controller

import (
	"github.com/gin-gonic/gin"
	uuidGen "github.com/google/uuid"
	"github.com/tot0p/SharePhoto/model"
	"github.com/tot0p/SharePhoto/utils/mongodb"
	"github.com/tot0p/SharePhoto/utils/session"
	"github.com/tot0p/env"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"path/filepath"
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

	bs, err := mongodb.DB.Find(env.Get("DATABASE_NAME"), "Collection", bson.M{
		"uuid": uuid,
	})

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(bs) == 0 {
		ctx.Redirect(302, "/"+uuid)
		return
	}

	pict := model.Picture{
		UUID:        uuidImg,
		Path:        "src/cdn/" + uuidImg + ext,
		AdderName:   "",
		Fingerprint: User.BrowserFingerPrinting,
		Like:        make([]string, 0),
		UUIDEvent:   uuid,
	}

	_, err = mongodb.DB.InsertOne(env.Get("DATABASE_NAME"), "Picture", pict)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(302, "/"+uuid)

}
