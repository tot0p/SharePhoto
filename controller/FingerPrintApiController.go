package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tot0p/SharePhoto/model"
	"github.com/tot0p/SharePhoto/utils/session"
	"net/http"
)

func FingerPrintApiController(ctx *gin.Context) {
	if session.SessionsManager.IsLogged(ctx) {
		ctx.JSON(http.StatusOK, gin.H{"success": "true"})
		return
	}
	// Récupérez les données de l'empreinte de navigateur depuis le corps de la requête
	var fingerprintData map[string]interface{}
	if err := ctx.ShouldBindJSON(&fingerprintData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fingerprint, ok := fingerprintData["fingerprint"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "fingerprint is not a string"})
		return
	}

	fmt.Println("Received Browser Fingerprint:", fingerprint)

	ip := ctx.ClientIP()

	fmt.Println("Received IP:", ip)

	User := model.User{
		Ip:                    ip,
		BrowserFingerPrinting: fingerprint,
		UUID:                  uuid.New().String(),
	}

	session.SessionsManager.CreateSession(ctx, &User)

	ctx.JSON(http.StatusOK, gin.H{"success": "true"})
}
