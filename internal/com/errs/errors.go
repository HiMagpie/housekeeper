package errs

import (
	"strings"
	"strconv"
	"runtime"
	"fmt"
	"housekeeper/internal/com/logger"
)

type ErrStr string

const (
	ERR_NONE = 0
	ERR_SEP = "=>"
	ERR_PARAM ErrStr = "1=>参数错误"
	ERR_OPERATION ErrStr = "2=>操作失败"

	// 配置相关
	CONF_FIELD_NOT_EXIST ErrStr = "101=>配置字段不存在"
	APP_ID_SECRET_REQUIRED ErrStr = "102=>APP_ID或APP_SECRET无效"
	CLIENT_KEY_INVALID ErrStr = "103=>无法识别客户端标示"

	// 缓存相关
	RC_PING_FAIL ErrStr = "201=>Redis无法Ping通"
	RC_REG_PUSH_SERVER_FAIL ErrStr = "202=>注册PushServer失败"

)

//返回错误码
func (errstr ErrStr) GetErrno() int {
	s := strings.Split(string(errstr), ERR_SEP)
	errno, _ := strconv.Atoi(s[0])
	return errno
}

//返回错误信息
func (err ErrStr) GetErrmsg() string {
	return strings.Split(string(err), ERR_SEP)[1]
}

// 实现Error interface
func (err ErrStr) Error() string{
	return string(err)
}

/**
 * 判断错误
 */
func CheckError(err error, data ...interface{}) bool{
	if err == nil {
		return false
	}
	logFilename := "proc"

	// 获取调用CheckError的方法名
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		f := runtime.FuncForPC(pc)
		path := strings.Split(f.Name(), "/")
		if len(path) > 0 {
			logFilename = path[len(path) - 1]
		}
	}

	// log记录
	logger.Error(logFilename, logger.Format("err", err.Error(), "occur", fmt.Sprintf("%s:%d", file, line), "data", data))
	return true
}
