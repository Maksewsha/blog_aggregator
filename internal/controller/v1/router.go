package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/maksewsha/blog_aggregator/internal/controller/v1/middleware"
	"github.com/maksewsha/blog_aggregator/pkg/log"
)

func NewRouter(logger log.MyLogger) *gin.Engine {
	router := gin.Default()

	rg := router.Group("/")
	rg.Use(middleware.MyLogger(logger))
	{
		newHomeHandlers(rg)
	}

	return router
}
