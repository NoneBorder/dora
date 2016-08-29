package doraemon

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

var Logger *logs.BeeLogger

func init() {
	Logger = logs.NewLogger(0)
	Logger.SetLogger("console", "")
	Logger.EnableFuncCallDepth(true)
}

func ResetLogger(buffer int64, provider string, params interface{}) (err error) {
	Logger = logs.NewLogger(buffer)

	paramsJ, err := json.Marshal(params)
	if err != nil {
		return
	}

	Logger.SetLogger(provider, string(paramsJ))
	Logger.EnableFuncCallDepth(true)
	return
}
