package workers

import (
	"time"
	"housekeeper/internal/cache"
	"housekeeper/internal/com/logger"
)

/**
 * 清理已经发送成功
 */
func ClearPushedMsgs() {
	go func() {
		for {
			msgIdCouters, err := cache.GetAllMsgIdCounters()
			if err != nil {
				logger.Error("workers.clear.pushed.msgs", logger.Format("err", err.Error()))
				time.Sleep(time.Second)
				continue
			}
			if len(msgIdCouters) == 0 {
				time.Sleep(time.Second)
				continue
			}

			logger.Debug("workers.clear.pushed.msgs", logger.Format("msg_id_couters", msgIdCouters))
			for msgId, counter := range msgIdCouters {
				totalNum, err := cache.GetMsgCidsNum(msgId)
				logger.Debug("workers.clear.pushed.msgs", logger.Format("msg_id", msgId, "total_num", totalNum, "cur_num", counter))
				if err != nil {
					continue
				}

				// 执行清除
				if totalNum <= counter {
					cache.DelMsgInfoAndCids(msgId)
					cache.DelMsgIdCounter(msgId)
				}
			}

			time.Sleep(time.Second * time.Duration(3))
		}
	}()
}

/**
 * 周期性地清理(按照消息体的时间进行判断消息是否过期)
 */
func ClearCyclicity() {

}