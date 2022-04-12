package model

type Result struct {
	Succeed bool        `json:"succeed"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
