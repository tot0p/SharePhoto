package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sharephoto/utils/mongodb"
	"sharephoto/utils/session"
)

func LikeController(ctx *gin.Context) {
	User := session.SessionsManager.GetUser(ctx)
	if User == nil {
		ctx.JSON(401, gin.H{
			"error": "not logged",
		})
		return
	}

	uuid := ctx.Param("uuid")
	uuidImg := ctx.Param("uuidImg")

	// get img info from db
	collection, err := mongodb.DB.Find(env.Get("DATABASE_NAME"), "Collection", bson.M{
		"uuid": uuid,
	})
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(collection) == 0 {
		ctx.JSON(404, gin.H{
			"error": "collection not found",
		})
		return
	}

	// get img from db

	fmt.Println("uuidImg", uuidImg)

	img, err := mongodb.DB.Find(env.Get("DATABASE_NAME"), "Picture", bson.M{
		"uuid": uuidImg,
	})
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(img) == 0 {
		ctx.JSON(404, gin.H{
			"error": "img not found",
		})
		return
	}

	var like2 = make([]string, 0)

	if l, ok := img[0]["like"].(primitive.A); !ok {
		img[0]["like"] = make([]string, 0)
	} else {
		if func() bool {
			for _, v := range l {
				if v.(string) == User.UUID {
					return true
				}
			}
			return false
		}() {
			for _, v := range l {
				if v.(string) != User.UUID {
					like2 = append(like2, v.(string))
				}
			}
		} else {
			like2 = append(like2, User.UUID)
			for _, v := range l {
				like2 = append(like2, v.(string))
			}
		}
	}

	_, err = mongodb.DB.UpdateOne(env.Get("DATABASE_NAME"), "Picture", bson.M{
		"uuid": uuidImg,
	}, bson.D{{Key: "$set", Value: bson.M{
		"like": like2,
	}}})

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(302, "/"+uuid)
}
