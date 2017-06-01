package controllers

import (
	"net/http"
	"encoding/json"
	"strings"
	"housekeeper/internal/com/logger"
	"housekeeper/internal/com/errs"
	"housekeeper/internal/model/third"
	"housekeeper/internal/cache/queue"
	"housekeeper/internal/model/db/dao"
	"housekeeper/internal/http_srv/uuid"
)

/**
 * receive multi messages from third-part app to push.
 */
func RecvMulti(res http.ResponseWriter, req *http.Request) {
	msgStr := req.FormValue("msg")
	towards := req.FormValue("towards")
	appId := req.FormValue("app_id")
	masterSecret := req.FormValue("master_secret")
	logger.Debug("recv.multi", logger.Format("msg_str", msgStr, "towards", towards, "app_id", appId, "master_secret", masterSecret))

	// 验证身份
	if exist, _ := dao.NewApp().ValidMaster(appId, masterSecret); !exist {
		res.Write(setFailData("", errs.ERR_PARAM))
		return
	}

	// 消息
	msg := third.NewMsg()
	err := json.Unmarshal([]byte(msgStr), msg)
	if err != nil {
		res.Write(setFailData(msgStr, errs.ERR_PARAM))
		return
	}

	// 验证 cid 所属
	cids := strings.Split(towards, ",")
	if len(cids) == 0 {
		res.Write(setFailData("", errs.ERR_PARAM))
		return
	}
	_, invalidCids, err := dao.NewAppCid().ValidCids(appId, cids)
	if len(invalidCids) > 0 {
		res.Write(setFailData(map[string][]string{
			"invalic_cids":  invalidCids,
		}, errs.ERR_PARAM))
		return
	}

	//msgId := cache.GenMsgId()
	msgId, err := uuid.GenMsgId()
	if err != nil {
		logger.Error("uuid", logger.Format("err", err.Error()))
		res.Write(setFailData("", errs.ERR_OPERATION))
		return
	}
	_, err = dao.NewMsgInfo().Create(msg, msgId, towards)
	if err != nil {
		res.Write(setFailData(err.Error(), errs.ERR_OPERATION));
		return
	}

	for _, cid := range cids {
		dao.NewMsgCidStatus().Create(cid, msgId)
	}

	logger.Info("recv.multi", logger.Format("msg_str", msgStr, "towards", towards, "app_id", appId, "master_secret", masterSecret))
	msgInfo, err := dao.NewMsgInfo().GetByMsgId(msgId)
	if err != nil {
		logger.Error("receiver", logger.Format("err", err.Error()))
		return
	}
	queue.PushMsgInfoToQueue(msgInfo)
	res.Write(setSuccData(map[string]int64{"msg_id": msgId}))
}
