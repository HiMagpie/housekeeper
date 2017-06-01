package queue

import (
	"strings"
	"fmt"
	"encoding/json"
	"housekeeper/internal/cache"
	"housekeeper/internal/model"
	"housekeeper/internal/com/logger"
	"housekeeper/internal/model/db/dao"
)

/**
 * Notify server (which connected with cid) to deal msg.
 */
func NotifyCidServer(cid string) {
	serverQueue, err := cache.GetServerNotifyQueueByCid(cid)
	if err != nil {
		logger.Error("notify.cid.server", logger.Format("err", err.Error() + "," + cid + ": no server queue to notify", "cid", cid, "queue", serverQueue))
		return
	}

	cache.PushToQueue(serverQueue, cid)
}

/**
 * Get cid's msg queue (`msg_id` waiting to push)
 */
func GetQueueByCid(cid string) string {
	return "mq_" + cid
}

/**
 * Deal MsgInfo, to separate msg body and msg_id
 * use msg_id to as identifier for saving memory
 */
func PushMsgInfoToQueue(msgInfo *dao.MsgInfo) error {
	// msg info ( means: msg body)
	err := setMsgEntity(msgInfo)
	if err != nil {
		logger.Error("push.msg.queue", logger.Format("err", err.Error(), "info", *msgInfo))
		return err
	}

	// msg's cid map (msg_id => cids)
	cids := strings.Split(msgInfo.Cids, ",")
	err = setMsgCids(msgInfo.MsgId, cids)
	if err != nil {
		logger.Error("push.msg.queue", logger.Format("err", err.Error(), "info", *msgInfo))
		return err
	}

	// push msg_id into cid's `waiting to push queue`
	logger.Info("cache.msg.pushmsginfo", logger.Format("msg_id", msgInfo.MsgId, "cids", cids))
	for _, cid := range cids {
		qname := GetQueueByCid(cid)
		_, err = cache.PushToQueue(qname, fmt.Sprintf("%d", msgInfo.MsgId))
		if err != nil {
			logger.Error("cache.msg.pushmsginfo", logger.Format("err", err.Error(), "queue", qname, "msg_id", msgInfo.MsgId, ))
			continue
		}

		// notify the server (which connected with cid client) to push msg
		NotifyCidServer(cid)
	}

	return nil
}

/**
 * Save msg body to map
 * e.g: 4 => `msg body json string`
 */
func setMsgEntity(msgInfo *dao.MsgInfo) error {
	key := cache.GetMsgInfoKey(msgInfo.MsgId)
	msgEntity := models.NewMsgEntity()
	msgEntity.MsgId = msgInfo.MsgId
	msgEntity.Ring = msgInfo.Ring
	msgEntity.Vibrate = msgInfo.Vibrate
	msgEntity.Cleanable = msgInfo.Cleanable
	msgEntity.Trans = msgInfo.Trans
	msgEntity.Begin = msgInfo.Begin
	msgEntity.End = msgInfo.End
	msgEntity.Title = msgInfo.Title
	msgEntity.Text = msgInfo.Text
	msgEntity.Logo = msgInfo.Logo
	msgEntity.Url = msgInfo.Url
	msgEntity.Ctime = msgInfo.Ctime
	infoBytes, err := json.Marshal(msgInfo)
	if err != nil {
		return err
	}
	return cache.HSet(key, fmt.Sprintf("%d", msgInfo.MsgId), string(infoBytes))
}

/**
 * map msg_id with msg's cids
 * for identify msg's pushing status.
 */
func setMsgCids(msgId int64, cids []string) error {
	key := cache.GetMsgCidsKey(msgId)
	return cache.SAdd(key, cids...)
}

