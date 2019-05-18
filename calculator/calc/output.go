package calc

type Output struct {
	Result float32 `json:"result" binding:"required" example:"4.6"`
}
