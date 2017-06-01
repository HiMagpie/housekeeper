package workers

import (
	"time"
	"gopkg.in/redis.v3"
	"housekeeper/internal/cache/queue"
	"housekeeper/internal/com/logger"
	"housekeeper/internal/cache"
	"housekeeper/internal/model/db/dao"
)

func DealMsgCidStatus() {
	go func() {
		stat := dao.NewMsgCidStatus()
		msgStat := queue.NewMsgStat()
		for {
			err := msgStat.BRPopFromQueue()
			if err != nil{
				if err != redis.Nil {
					logger.Error("workers.update.msg.status", logger.Format("err", err.Error()))
				}
				time.Sleep(time.Second)
				continue
			}

			if msgStat.Status == queue.MSG_STAT_PUSH_SUCC {
				cache.IncMsgIdCounter(msgStat.MsgId)
			}

			logger.Debug("workers.update.msg.status", logger.Format("cid", msgStat.Cid, "status", msgStat.Status))
			err = stat.UpdateStatusByCidMsgId(msgStat.Cid, msgStat.MsgId, msgStat.Status)
			if err != nil {
				logger.Error("workers.update.msg.status", logger.Format("err", err.Error(), "cid", msgStat.Cid, "status", msgStat.Status))
			}
		}
	}()
}
