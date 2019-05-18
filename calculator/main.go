package main

import (
	"crypto/tls"
	"net/http"
	"time"

	"samples/calculator/calc"
	"samples/calculator/cfg"
	"samples/calculator/log"
	"samples/calculator/ping"

	_ "samples/calculator/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	logger   log.Logger
	calcCtrl calc.CalculatorController
	pingCtrl ping.PingController
)

// @title Swagger Samples/Calculator API
// @version 1.0
// @description This service provides four main mathematical operations
// @termsOfService http://www.samples.io/terms/

// @contact.name Samples API Support
// @contact.url http://www.samples.io/support
// @contact.email support@samples.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /
func main() {
	logger = log.NewLogger()
	port := configServices()

	logger.Log("Starting to create engin...")
	engine := gin.Default()
	engine.POST("/api/calc/:op", calcCtrl.Calculate)
	engine.GET("/ping", pingCtrl.Ping)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Log("Starting to run application on port %s...", port)
	engine.Run(port)
}

func configServices() string {
	configSvc := cfg.NewConfigService()
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

	calcSvc := calc.NewCalculatorService()
	calcCtrl = calc.NewCalculatorController(calcSvc, logger)
	pingCtrl = ping.NewPingController(client, logger)

	return config.Port
}
