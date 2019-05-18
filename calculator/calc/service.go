package calc

import (
	"errors"
)

type CalculatorService interface {
	Add(input Input) (Output, error)
	Sub(input Input) (Output, error)
	Mul(input Input) (Output, error)
	Div(input Input) (Output, error)
}

type calculatorService struct {
}

func NewCalculatorService() CalculatorService {
	return &calculatorService{}
}

func (svc *calculatorService) Add(input Input) (Output, error) {
	result := input.FirstNumber + input.SecondNumber
	return Output{
		Result: result,
	}, nil
}

func (svc *calculatorService) Sub(input Input) (Output, error) {
	result := input.FirstNumber - input.SecondNumber
	return Output{
		Result: result,
	}, nil
}

func (svc *calculatorService) Mul(input Input) (Output, error) {
	result := input.FirstNumber * input.SecondNumber
	return Output{
		Result: result,
	}, nil
}

func (svc *calculatorService) Div(input Input) (Output, error) {
	if input.SecondNumber == 0 {
		return Output{}, errors.New("Divid by zero error")
	}
	result := input.FirstNumber / input.SecondNumber
	return Output{
		Result: result,
	}, nil
}
