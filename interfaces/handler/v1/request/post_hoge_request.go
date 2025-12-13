package request

type PostHogeRequestBody struct {
	HogeName string `json:"hoge_name" binding:"required"`
}
