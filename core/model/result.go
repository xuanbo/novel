package model

type Result struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func Ok(data interface{}) *Result {
	return &Result{200, "OK", data}
}

func Fail(data interface{}) *Result {
	return &Result{500, "Fail", data}
}

func NotFound() *Result {
	return &Result{404, "Not Found", ""}
}