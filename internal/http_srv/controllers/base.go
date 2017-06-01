package controllers
import (
	"encoding/json"
	"housekeeper/internal/com/errs"
)

func formatRet(data interface{}, code interface{}, msg interface{}) map[string]interface{} {
	ret := map[string]interface{} {
		"msg": msg,
		"code": code,
		"data": data,
	}

	return ret
}

/**
 * 成功返回的数据格式
 * @TODO 加密数据流
 */
func setSuccData(data interface{}) []byte {
	rb, _ := json.Marshal(formatRet(data, 0, "ok"))
	return rb
}

/**
 * 失败时,返回的数据格式
 * @TODO 加密数据流
 * @param data
 * @param code
 * @param msg
 */
func setFailData(data interface{}, err errs.ErrStr) []byte {
	if data == nil {
		data = ""
	}
	rb, _ := json.Marshal(formatRet(data, err.GetErrno(), err.GetErrmsg()))
	return rb
}
