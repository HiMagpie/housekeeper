package queue

import (
	"encoding/json"
	"time"
	"gopkg.in/redis.v3"
	"housekeeper/internal/cache"
	"housekeeper/internal/com/logger"
)

const (
	KEY_QUEUE_CID_VALID = "queue_valid_cid"
)

type CidValid struct {
	Cid      string `json:"cid"`
	Hostname string `json:"hostname"`
}

func NewCidValid() *CidValid {
	return new(CidValid)
}

func (this *CidValid) BRPopFromQueue() error {
	item, err := cache.BRPopFromQueue(KEY_QUEUE_CID_VALID, time.Second * time.Duration(3))
	if err != nil {
		goto err_bpop_cid_valid
	}

	logger.Debug("queue.cid_valid.bpop.fromqueue", logger.Format("item", item))
	err = json.Unmarshal([]byte(item), this)
	if err != nil {
		goto err_bpop_cid_valid
	}

	// 操作失败
	err_bpop_cid_valid:
	if err != nil {
		if err != redis.Nil {
			logger.Error("queue.cid_valid.bpop.fromqueue", logger.Format("err", err.Error(), "item", item))
		}

		return err
	}

	return nil
}
