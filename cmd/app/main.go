package main

import (
	"github.com/joho/godotenv"
	"github.com/maksewsha/blog_aggregator/config"
	"github.com/maksewsha/blog_aggregator/internal/app"
	"github.com/maksewsha/blog_aggregator/pkg/log"
)

func main() {
	godotenv.Load()
	cfg := config.NewConfig()
	app := app.NewApp(cfg, log.NewAppLogger())
	app.Start()
}
