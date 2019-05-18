package calc

import (
	"errors"
	"fmt"
	"net/http"
	"samples/calculator/log"
	"strings"

	"github.com/gin-gonic/gin"
)

type CalculatorController interface {
	Calculate(ctx *gin.Context)
}

type calculatorController struct {
	calcSvc CalculatorService
	logger  log.Logger
}

func NewCalculatorController(calcSvc CalculatorService, logger log.Logger) CalculatorController {
	return &calculatorController{
		calcSvc: calcSvc,
		logger:  logger,
	}
}

// Calculate godoc
// @Summary Calculate runs requested operation
// @Description Calculate runs requested operation on input coming in request body
// @ID Calculate
// @Accept  json
// @Produce  json
// @Param op path string true "Operation add|sub|mul|div"
// @Param input body calc.Input true "Json representaion of an Input object"
// @Success 200 {object} calc.Output
// @Failure 400 string string "Provided op xyz is not valid"
// @Failure 500 string string "Internal Server Error"
// @Router /api/calc/{op} [get]
func (ctrl *calculatorController) Calculate(ctx *gin.Context) {
	op := strings.ToLower(ctx.Params.ByName("op"))

	var input Input
	err := ctx.BindJSON(&input)
	if err != nil {
		ctrl.logger.Error(errors.New("Could not bind incoming JSON"))
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var results Output
	ctrl.logger.Log("Preparing sending calculation request...")
	switch op {
	case "add":
		results, err = ctrl.calcSvc.Add(input)
	case "sub":
		results, err = ctrl.calcSvc.Sub(input)
	case "mul":
		results, err = ctrl.calcSvc.Mul(input)
	case "div":
		results, err = ctrl.calcSvc.Div(input)
	default:
		err := fmt.Errorf("Provided op %s is not valid", op)
		ctrl.logger.Error(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err != nil {
		ctrl.logger.Error(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctrl.logger.Log("Results came successfully %v", results.Result)
	ctx.JSON(http.StatusOK, results)
}
