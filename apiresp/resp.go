package apiresp

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const NbReqStatusHeader = "Nb-Req-Status"

type ApiResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewResp(d interface{}, msgs ...string) *ApiResp {
	var msg string
	if len(msgs) > 0 {
		msg = msgs[0]
	}

	return &ApiResp{
		Code: 200,
		Msg:  msg,
		Data: d,
	}
}

func NewErr(msg string, ds ...interface{}) *ApiResp {
	var d interface{}
	if len(ds) > 0 {
		d = ds[0]
	}

	return &ApiResp{
		Code: 207,
		Msg:  msg,
		Data: d,
	}
}

func NewDetail(code int, msg string, ds ...interface{}) *ApiResp {
	var d interface{}
	if len(ds) > 0 {
		d = ds[0]
	}

	return &ApiResp{
		Code: code,
		Msg:  msg,
		Data: d,
	}
}

func (self *ApiResp) IsSuccess() bool {
	return self.Code == 200
}

// BeegoServeJSON 返回 json 内容
// @deprecated
func (self *ApiResp) BeegoServeJSON(c beego.Controller) {
	c.Ctx.Output.SetStatus(self.Code)
	c.Data["json"] = self
	c.ServeJSON()
	c.StopRun()
}

// ReturnJSON 返回 json 内容
// httpCode 默认为200， 业务请求状态码在 response header 中的 "nb-req-status" 体现
func (self *ApiResp) ReturnJSON(c beego.Controller, httpCode ...int) {
	httpCode = append(httpCode, 200)
	c.Ctx.Output.SetStatus(httpCode[0])
	c.Ctx.Output.Header(NbReqStatusHeader, strconv.Itoa(self.Code))
	c.Data["json"] = self
	c.ServeJSON()
	c.StopRun()
}

func (self *ApiResp) JSON(ctx *context.Context, httpCode ...int) {
	httpCode = append(httpCode, 200)
	ctx.Output.SetStatus(httpCode[0])
	ctx.Output.Header(NbReqStatusHeader, strconv.Itoa(self.Code))
	ctx.Output.JSON(self, false, false)
	panic(beego.ErrAbort)
}
