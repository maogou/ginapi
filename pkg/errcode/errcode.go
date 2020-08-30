package errcode

import (
	"fmt"
	"net/http"
)

//错误返回的结构体
type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

//检测业务码是否已经存在
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在,请更换一个", code))
	}

	codes[code] = msg

	return &Error{
		code: code,
		msg:  msg,
	}
}

//获取错误信息
func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.code, e.msg)
}

//获取错误码
func (e *Error) Code() int {
	return e.code
}

//获取错误信息
func (e *Error) Msg() string {
	return e.msg
}

//格式化错误信息
func (e *Error)Msgf(args []interface{}) string  {
	return fmt.Sprintf(e.msg,args...)
}

//获取错误详细信息
func (e *Error) Details() []string  {
	return e.details
}

//追加详细错误信息
func (e *Error)WithDetail(detail ...string) *Error  {
	e.details = []string{}
	for _,d := range detail{
		e.details = append(e.details,d)
	}

	return e
}

//返回对具体的状态码进行转换
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
