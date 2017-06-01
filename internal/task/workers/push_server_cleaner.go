package workers

import (
	"time"
	"housekeeper/internal/com/cfg"
	"housekeeper/internal/cache"
)

func CleanPushServer() {
	go func() {
		ps := cache.NewPushServerCache()
		for {
			ps.CleanInvalidPushServer()
			time.Sleep(time.Second * time.Duration(cfg.C.RpcSrv.PushServerHb))
		}
	}()
}
