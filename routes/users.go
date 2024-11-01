package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": fmt.Sprintf("%v", http.StatusBadRequest), "message": "Could not parse request data"})
		return
	}

	user.Id = 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	lastId, err := user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": fmt.Sprintf("%v", http.StatusInternalServerError), "message": "Could not save user to database"})
		return
	}

	confirmation := fmt.Sprintf("Successfully created user with ID: %v", lastId)

	context.JSON(http.StatusOK, gin.H{"message": confirmation})
}

func getAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve all users", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}

func getUser(context *gin.Context) {
	rawId := context.Param("id")
	userId, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	user, err := models.GetUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user details"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse your request body"})
		return
	}

	err = user.LoginUser()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authorize user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
