package controllers

import (
	"net/http"

	"samples/spotify-proxy/services"

	"github.com/gin-gonic/gin"
)

type PingController interface {
	Ping(ctx *gin.Context)
}

type pingController struct {
	client *http.Client
	logger services.Logger
}

func NewPingController(client *http.Client, logger services.Logger) PingController {
	return &pingController{
		client: client,
		logger: logger,
	}
}

func (ctrl *pingController) Ping(ctx *gin.Context) {
	ctrl.logger.Log("Start pinging...")
	if ok, err := canPingGoogle(ctrl.client); ok {
		ctx.String(http.StatusOK, "pong")
		ctrl.logger.Log("Finished pinging successfully")
	} else {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		ctrl.logger.Error(err)
	}
}

func canPingGoogle(client *http.Client) (bool, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		return false, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == http.StatusOK, nil
}
