package calc

type Input struct {
	FirstNumber  float32 `json:"firstNumber" binding:"required" example:"1.2"`
	SecondNumber float32 `json:"secondNumber" binding:"required" example:"3.4"`
}
