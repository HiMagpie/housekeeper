package uuid

import (
	"github.com/zheng-ji/goSnowFlake"
	"housekeeper/internal/com/cfg"
)

var (
	iw *goSnowFlake.IdWorker
)

func GenMsgId() (int64, error) {
	var err error
	iw, err = goSnowFlake.NewIdWorker(int64(cfg.C.HttpSrv.SnowFlakeWorkerId))
	if err != nil {
		return 0, err
	}

	return iw.NextId()
}

