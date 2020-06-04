package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"hulujia/util/log"
	"hulujia/util/response"
	"reflect"
)

type BaseController struct {
}

// 表单验证
func (c *BaseController) Validate(ctx *gin.Context, obj interface{}) bool {
	valid := validation.Validation{}
	b, err := valid.Valid(obj)
	if err != nil {
		log.Error(err.Error())
		c.HttpFailed(ctx, 500,err.Error())
		return false
	}
	if !b {
		st := reflect.TypeOf(obj)
		filed, _ := st.FieldByName(valid.Errors[0].Field)
		var alias = filed.Tag.Get("alias")
		c.ValidFailed(ctx, alias + " " + valid.Errors[0].Message)
		return false
	}
	return true
}

func (c *BaseController) Message(ctx *gin.Context, message string) {
	var req = response.Response{ctx}
	req.Message(message)
}


func (c *BaseController) Success(ctx *gin.Context, data interface{}) {
	var req = response.Response{ctx}
	req.Success(data)
}

func (c *BaseController) SuccessWithMsg(ctx *gin.Context, msg string, data interface{}) {
	var req = response.Response{ctx}
	req.SuccessWithMsg(msg, data)
}

func (c *BaseController) Failed(ctx *gin.Context, errCode int) {
	var req = response.Response{ctx}
	req.Failed(errCode)
	return
}

func (c *BaseController) FailedWithMsg(ctx *gin.Context, errCode int, message string) {
	var req = response.Response{ctx}
	req.FailedWithMsg(errCode,message)
}

func (c *BaseController) ValidFailed(ctx *gin.Context, message string) {
	var req = response.Response{ctx}
	req.ValidFailed(message)
}

func (c *BaseController) HttpFailed(ctx *gin.Context, httpCode int, message string) {
	var req = response.Response{ctx}
	req.HttpFailed(httpCode, message)
}


