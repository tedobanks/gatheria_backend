package middlewares

import (
	"fmt"
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

// This file handles Middleware, Middleware are basically functions that execute before the main function handler. In this case it can help save our overhead cost on a request. It checks if the user is authenticated first then is stops the execution at this point. if the user is authenticated it lets them get through to the next handler.

// context.Set() allows us to move data from one function to another provoded that they use the same context

func Authenicate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": fmt.Sprintf("%v", http.StatusUnauthorized), "message": "Please pass a token to access this route"})
		return
	}

	id, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token: " + err.Error()})
		return
	}

	context.Set("id", id)
	context.Next()
}
