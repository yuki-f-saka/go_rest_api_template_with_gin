package request

type GetHogeRequestHeader struct {
	XRequestID string `header:"X-Request-ID" binding:"required"`
	HogeID     string `uri:"hoge_id" binding:"required"`
}
