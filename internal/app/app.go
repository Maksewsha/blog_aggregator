package app

import (
	"github.com/gin-gonic/gin"
	"github.com/maksewsha/blog_aggregator/config"
	"github.com/maksewsha/blog_aggregator/internal/controller/v1"
	"github.com/maksewsha/blog_aggregator/pkg/log"
)

type app struct {
	config *config.Config
	logger log.MyLogger
	router *gin.Engine
}

func NewApp(cfg *config.Config, logger log.MyLogger) *app {
	router := v1.NewRouter()

	return &app{
		config: cfg,
		logger: logger,
		router: router,
	}
}

func (a *app) Start() {
	err := a.router.Run(a.config.HOST + ":" + a.config.PORT)
	if err != nil {
		a.logger.Error(err.Error())
	}
}
