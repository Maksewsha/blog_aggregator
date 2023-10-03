package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maksewsha/blog_aggregator/pkg/log"
)

func MyLogger(logger log.MyLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info(fmt.Sprintf("Received request %s %s", c.Request.Method, c.FullPath()))
		c.Next()
		logger.Info(fmt.Sprintf("Response to request with %d", c.Writer.Status()))
	}
}
