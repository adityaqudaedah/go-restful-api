package web

type WebResponse struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
	Data    interface{} `json:"data"`
}