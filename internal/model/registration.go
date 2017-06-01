package models

import (
	"fmt"
	"housekeeper/internal/com/utils"
)

/**
 * 验证通过, 返回cid等信息
 */
type ValidatedBack struct {
	Cid     string `json:"cid"`
	Servers []string `json:"servers"`
}

// @TODO 解密checkIn传过来的数据
func DecryptRegData(data []byte) (appUuid, os, appVer string, err error) {

	return
}

// 生成cid
// @TODO 校验Cid唯一性
func GenCid(appId, appSecret, clientKey string) string {
	return utils.Md5Str(fmt.Sprintf("%s%s%sKiLh0sA", appId, appSecret, clientKey))
}

// @TODO 通过cid获取长连接的服务器
func GetLongServerByCid(cid string) []string {
	servers := []string{}
	servers = append(servers, "127.0.0.1:7777")
	servers = append(servers, "127.0.0.1:7777")
	servers = append(servers, "127.0.0.1:7777")
	return servers
}
