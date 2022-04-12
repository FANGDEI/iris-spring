package handlers

type Result struct {
	Succeed bool        `json:"succeed"`
	Msg     string      `json:"msg"`
	Result  interface{} `json:"result"`
}
