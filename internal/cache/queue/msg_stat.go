package queue

import (
	"encoding/json"
	"errors"
	"time"
	"housekeeper/internal/cache"
	"housekeeper/internal/com/logger"
)

const (
	KEY_MSG_STATUS_QUEUE = "queue_msg_status"

	MSG_STAT_PUSH_ACK = 3 // 刚发送, 等待ACK
	MSG_STAT_PUSH_SUCC = 4 // 发送成功(收到消息ACK)
	MSG_STAT_OFFLINE = 5 // 客户端不在线, 进入离线状态
	MSG_STAT_OVERDUE = 6 // 过期失效
	MSG_STAT_FAIL = 7 // 失败
)

type MsgStat struct {
	MsgId  int64 `json:"msg_id"`
	Cid    string `json:"cid"`
	Status int `json:"status"`
}

func NewMsgStat() *MsgStat {
	return new(MsgStat)
}

func (this *MsgStat) BRPopFromQueue() error {
	res, err := cache.BRPopFromQueue(KEY_MSG_STATUS_QUEUE, time.Second * time.Duration(3))
	if err != nil {
		return err
	}
	if res == "" {
		logger.Debug("queue.msg_stat.popfromqueue", "-")
		return errors.New("empty queue.")
	}

	err = json.Unmarshal([]byte(res), this)
	if err != nil {
		logger.Error("queue.msg_stat.popfromqueue", logger.Format("err", err.Error()))
		return err
	}

	return nil
}
