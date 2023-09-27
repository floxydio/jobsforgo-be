package entities

type ErrorModel struct {
	StatusCode int    `json:"status"`
	ErrMessage string `json:"errMessage"`
	Message    string `json:"message"`
}
