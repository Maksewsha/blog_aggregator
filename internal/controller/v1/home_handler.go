package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func newHomeHandlers(sg *gin.RouterGroup) {
	sg.GET("/", homePage)
}

func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello at home page!",
	})
}
