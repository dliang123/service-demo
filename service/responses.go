package service

type CommonResponse struct {
	Code  int                      `json:"code"`
	Msg   string                   `json:"msg"`
	Data  map[string]interface{}   `json:"data"`
	Err   string                    `json:"err,omitempty"`
}


type SayHelloResponse struct {
	Result   string `json:"result"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}