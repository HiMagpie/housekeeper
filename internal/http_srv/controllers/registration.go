package controllers

import (
	"net/http"
	"housekeeper/internal/com/logger"
	"housekeeper/internal/model"
	"housekeeper/internal/com/errs"
	"housekeeper/internal/model/db/dao"
)

/**
 * 客户端check in, 获取cid和连接的server
 * @param app_id 应用的id
 * @param app_secret 应用的秘钥
 * @param client_key 客户端生成的唯一标示
 * @param ver 当前SDK的版本
 * @param os 系统版本
 * @param install_time 安装时间
 */
func CheckIn(res http.ResponseWriter, req *http.Request) {
	// 1.1 验证参数 @TODO Aes加密后的串
	appId := req.FormValue("app_id")
	appSecret := req.FormValue("app_secret")
	clientKey := req.FormValue("client_key")
	ver := req.FormValue("ver")
	os := req.FormValue("os")
	installTime := req.FormValue("install_time")

	flag, err := validParams(appId, appSecret, clientKey)
	if !flag {
		logger.Error("cntl.registration.checkin", logger.Format(
			"err", "valid params fail.",
			"app_id", appId, "app_secret", appSecret, "client_key", clientKey,
			"ver", ver, "os", os, "install_time", installTime,
		))

		res.Write(setFailData(nil, err))
		return
	}

	// 2.1 生成返回数据
	cid := models.GenCid(appId, appSecret, clientKey)
	dao.NewAppCid().CreateUnique(cid, appId)
	ret := &models.ValidatedBack{
		Cid: cid,
		Servers: models.GetLongServerByCid(cid),
	}

	res.Write(setSuccData(ret))
}

func validParams(appId, appSecret, clientKey string) (bool, errs.ErrStr) {
	if appId == "" || appSecret == "" {
		return false, errs.APP_ID_SECRET_REQUIRED
	}

	if clientKey == "" {
		return false, errs.CLIENT_KEY_INVALID
	}

	// 1.2 验证app_id&app_secret是否存在
	if exist, _ := dao.NewApp().ValidApp(appId, appSecret); !exist {
		return false, errs.APP_ID_SECRET_REQUIRED
	}

	return true, ""
}

/**
 * 客户端check in, 获取cid和连接的server
 * @param app_id 应用的id
 * @param app_secret 应用的秘钥
 * @param client_key 客户端生成的唯一标示
 * @param cid
 */
func GetServers(res http.ResponseWriter, req *http.Request) {
	// 1.1 验证参数 @TODO Aes加密后的串
	appId := req.FormValue("app_id")
	appSecret := req.FormValue("app_secret")
	clientKey := req.FormValue("client_key")
	cid := req.FormValue("cid")

	flag, err := validParams(appId, appSecret, clientKey)
	if !flag {
		logger.Error("cntl.registration.getservers", logger.Format(
			"app_id", appId, "app_secret",
			appSecret, "client_key", clientKey, "cid", cid,
		))

		res.Write(setFailData(nil, err))
		return
	}

	// 2.1 生成返回数据
	tmpCid := models.GenCid(appId, appSecret, clientKey)
	if cid != tmpCid {
		logger.Error("cntl.registration.getservers", logger.Format("cid", cid, "tmp_cid", tmpCid))
		res.Write(setFailData(nil, errs.APP_ID_SECRET_REQUIRED))
		return
	}
	ret := &models.ValidatedBack{
		Cid: cid,
		Servers: models.GetLongServerByCid(cid),
	}

	res.Write(setSuccData(ret))
}
