package controllers

import (
	"jwt-gin/pkg/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var reg UserInfo

	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registered"})
	// store RegInfo in the database
}

func Login(c *gin.Context) {

	var reg UserInfo

	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := LoginCheck(reg.Username, reg.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func LoginCheck(username string, password string) (string, error) {

	// verify username and password ...
	// if ok, will return token to client

	token, err := token.GenerateToken()

	if err != nil {
		return "", err
	}

	return token, nil
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": "hello"})
}
