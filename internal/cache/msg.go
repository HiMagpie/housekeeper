package cache

import (
	"fmt"
	"housekeeper/internal/com/utils"
	"housekeeper/internal/com/logger"
)

const (
	KEY_MAX_MSG_ID = "set_max_msg_id"
)

func GetMsgCidsNum(msgId int64) (int64, error) {
	c := rc.SCard(GetMsgCidsKey(msgId))
	return c.Result()
}

func GetMsgInfoKey(msgId int64) string {
	return "msg_info_" + utils.Md5AndSub(fmt.Sprintf("%d", msgId), 0, 4)
}

func GetMsgCidsKey(msgId int64) string {
	return fmt.Sprintf("msg_cids_%d", msgId)
}

/**
 * 删除msg的消息体和对应的cid Set
 */
func DelMsgInfoAndCids(msgId int64) error {
	msgInfoKey := GetMsgInfoKey(msgId)
	msgCidKey := GetMsgCidsKey(msgId)

	c := rc.HDel(msgInfoKey, fmt.Sprintf("%d", msgId))
	if c.Err() != nil {
		return c.Err()
	}
	c = rc.Del(msgCidKey)
	if c.Err() != nil {
		return c.Err()
	}

	return nil
}

/**
 * 生成msg_id
 */
func GenMsgId() int64 {
	c := rc.IncrBy(KEY_MAX_MSG_ID, 1)
	msgId, err := c.Result()
	if err != nil {
		logger.Error("cache.msg.genmsgid", logger.Format("err", err.Error(), "msg_id", msgId, ))
	}

	return msgId
}
