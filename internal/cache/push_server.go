package cache

import (
	"housekeeper/internal/com/logger"
	"fmt"
	"time"
	"gopkg.in/redis.v3"
	"housekeeper/internal/com/cfg"
)

const (
	Z_PUSH_SERVERS = "sorted_set_push_server"
)

type PushServer struct{}

func NewPushServerCache() *PushServer {
	return new(PushServer)
}

func (this *PushServer) AddUnique(ip, port string) error {
	v := fmt.Sprintf("%s:%s", ip, port)
	m := redis.Z{
		Score: float64(time.Now().Unix()),
		Member: v,
	}
	c := rc.ZAdd(Z_PUSH_SERVERS, m)

	if c.Err() != nil {
		logger.Error("cache.push_server.add", logger.Format("err", c.Err().Error(), "ip", ip, "port", port))
		return c.Err()
	}

	return nil
}

func (this *PushServer) CleanInvalidPushServer() (int64, error) {
	// push server info will be clean when heartbeat doesn't sent after three circles
	c := rc.ZRemRangeByScore(Z_PUSH_SERVERS, "0", fmt.Sprintf("%d", time.Now().Unix() - int64(cfg.C.RpcSrv.PushServerHb * 3)))

	if c.Err() != nil {
		logger.Error("push.server.cache", logger.Format("err", c.Err().Error(), "cmd", c.String()))
	}
	return c.Result()
}

