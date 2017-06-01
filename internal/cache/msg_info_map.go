package cache

import (
	"fmt"
	"housekeeper/internal/com/utils"
)

const (
	KEY_MSG_FIN_COUNTER = "hmap_msg_fin_counter"
)

/**
 * msg已经完成推送的cid计数
 * 用于计算是否可以清除掉消息体和消息对应的cid
 */
func IncMsgIdCounter(msgId int64) error {
	c := rc.HIncrBy(KEY_MSG_FIN_COUNTER, fmt.Sprintf("%d", msgId), 1);
	return c.Err()
}

/**
 * 删除msg的计数器
 */
func DelMsgIdCounter(msgId int64) error {
	c := rc.HDel(KEY_MSG_FIN_COUNTER, fmt.Sprintf("%d", msgId))
	return c.Err()
}

/**
 * msg_id => counter
 */
func GetAllMsgIdCounters() (map[int64]int64, error) {
	c := rc.HGetAll(KEY_MSG_FIN_COUNTER)
	if c.Err() != nil {
		return nil, c.Err()
	}

	// 处理
	ret := make(map[int64]int64)
	for i, v := range c.Val() {
		if i % 2 == 1 {
			msgId, err := utils.AtoInt64(c.Val()[i - 1])
			counter, err := utils.AtoInt64(v)
			if err != nil {
				continue
			}
			ret[msgId] = counter
		}
	}

	return ret, nil
}