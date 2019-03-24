package model

import "github.com/jinzhu/gorm"

//返回的模型

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (R *Result) Err(msg string) *Result {
	return &Result{0, msg, nil}
}

func (R *Result) Suc(msg string) *Result {
	return &Result{1, msg, nil}
}
func (R *Result) Get(code int, msg string, data interface{}) *Result {
	return &Result{code, msg, data}
}

//user
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Tel      string `json:"tel"`
}
