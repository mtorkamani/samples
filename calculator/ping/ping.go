package ping

import (
	"net/http"

	"samples/calculator/log"

	"github.com/gin-gonic/gin"
)

type PingController interface {
	Ping(ctx *gin.Context)
}

type pingController struct {
	client *http.Client
	logger log.Logger
}

func NewPingController(client *http.Client, logger log.Logger) PingController {
	return &pingController{
		client: client,
		logger: logger,
	}
}

// Ping godoc
// @Summary Ping the service and get 200 Ok response to make sure application is up and running
// @Description It is the health check endpoint
// @ID Ping
// @Produce  plain
// @Success 200 string string "Pong"
// @Failure 500 string string "Internal Server Error"
// @Router /ping [get]
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
