package workers

import (
	"time"
	"gopkg.in/redis.v3"
	"housekeeper/internal/cache/queue"
	"housekeeper/internal/com/logger"
)

const (
	VALID_FLAG_SUCC int32 = 2 // 验证成功
	VALID_FLAT_FAIL int32 = 3 // 验证失败
)

func ValidCidTask() {
	cidValid := queue.NewCidValid()
	cidValidRes := queue.NewCidValidRes()

	for {
		err := cidValid.BRPopFromQueue()
		if err != nil {
			if err != redis.Nil {
				logger.Error("workers.validcidtask", logger.Format("err", err.Error()))
			}
			time.Sleep(time.Second)

			continue
		}

		cidValidRes.Cid = cidValid.Cid
		cidValidRes.Valid = VALID_FLAG_SUCC
		cidValidRes.PushToQueue(cidValid.Hostname)
		logger.Debug("workers.validcidtask", "cid", cidValid.Cid, "res", VALID_FLAG_SUCC)
	}
}
