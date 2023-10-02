package app

import (
	"github.com/gin-gonic/gin"
	"github.com/maksewsha/blog_aggregator/cfg"
)

type App struct {
	Config *cfg.Config

	router *gin.Engine
}

func (a *App) Start(host, port string) {
	a.router = gin.Default()
	a.router.Group("/")
	{
		a.router.GET("/")
	}

	err := a.router.Run(host + ":" + port)
	if err != nil {

	}
}
