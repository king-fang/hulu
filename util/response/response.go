package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Gin *gin.Context
}

type ResponseData struct {
	Code int          `json:"code"`
	Msg string 		  `json:"msg"`
	Data interface{}  `json:"data"`
}

type ResponseMessage struct {
	Code int        `json:"code"`
	Msg string  	`json:"msg"`
}

// 成功消息返回
func (r *Response) Message(msg string)  {
	r.Gin.JSON(http.StatusOK, ResponseMessage{
		Code: http.StatusOK,
		Msg:  msg,
	})
	return
}

// 成功数据返回
func (r *Response) Success(data interface{})  {
	r.Gin.JSON(http.StatusOK, ResponseData{
		Code: http.StatusOK,
		Msg: "success",
		Data: data,
	})
	return
}

// 返回数据带消息
func (r *Response) SuccessWithMsg(msg string,data interface{})  {
	r.Gin.JSON(http.StatusOK,ResponseData{
		Code: http.StatusOK,
		Msg: msg,
		Data: data,
	})
	return
}

// 失败返回
func (r *Response) Failed(errCode int)  {
	r.Gin.JSON(http.StatusOK, ResponseMessage{
		Code: errCode,
		Msg:  GetMsg(errCode,""),
	})
	return
}

// 失败返回带自定义消息
func (r *Response) FailedWithMsg(errCode int, message string)  {
	r.Gin.JSON(http.StatusOK,ResponseMessage{
		Code: errCode,
		Msg:  GetMsg(errCode,message),
	})
	return
}

// http状态码返回
func (r *Response) HttpFailed(httpCode int, message string)  {
	r.Gin.JSON(httpCode, ResponseMessage{
		Code: httpCode,
		Msg:  message,
	})
	return
}

// 数据验证失败返回
func (r *Response) ValidFailed(message string)  {
	r.HttpFailed(http.StatusUnprocessableEntity,message)
	return
}