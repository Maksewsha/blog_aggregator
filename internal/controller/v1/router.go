package v1

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()

	rg := router.Group("/")
	{
		newHomeHandlers(rg)
	}

	return router
}
