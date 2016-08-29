package apiresp

import "github.com/astaxie/beego"

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

func (self *ApiResp) BeegoServeJSON(c beego.Controller) {
	c.Ctx.Output.SetStatus(self.Code)
	c.Data["json"] = self
	c.ServeJSON()
}
