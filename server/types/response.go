package types

type Response struct {
	Message    string      `json:"message"`
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}
