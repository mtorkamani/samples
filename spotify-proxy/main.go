package main

import (
	"crypto/tls"
	"net/http"
	"time"

	"samples/spotify-proxy/controllers"
	"samples/spotify-proxy/services"

	"github.com/gin-gonic/gin"
)

var (
	logger     services.Logger
	searchCtrl controllers.SearchController
	pingCtrl   controllers.PingController
)

func main() {
	logger = services.NewLogger()

	port := configServices()

	logger.Log("Starting to create engin...")
	engine := gin.Default()
	engine.GET("/api/search/:term", searchCtrl.Search)
	engine.GET("/ping", pingCtrl.Ping)

	logger.Log("Starting to run application on port %s...", port)
	engine.Run(port)
}

func configServices() string {
	configSvc := services.NewConfigService(logger)
	config, err := configSvc.Get()
	if err != nil {
		logger.Fatal(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Timeout:   time.Second * 20,
		Transport: tr,
	}

	authSvc := services.NewAuthService(logger)
	searchSvc := services.NewSearchService(logger)

	searchCtrl = controllers.NewSearchController(client, authSvc, searchSvc, config, logger)
	pingCtrl = controllers.NewPingController(client, logger)

	return config.Port
}
