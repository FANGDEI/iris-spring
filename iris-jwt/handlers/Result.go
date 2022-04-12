package handlers

type Result struct {
	Succeed bool        `json:"succeed"`
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
}
