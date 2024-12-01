package middlewares

import (
	"log"
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authentiate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not anthorized."})
		return
	}

	userId, err := utils.VerifyToken(token)
	log.Println(err)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not anthorized. "})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
