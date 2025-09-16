package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func UpdateProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateProfile handler - to be implemented",
	})
}

func GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetProfile handler - to be implemented",
	})
}
