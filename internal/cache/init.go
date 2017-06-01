package cache

import (
	"fmt"
	"strings"
	"gopkg.in/redis.v3"
	"housekeeper/internal/com/cfg"
	"housekeeper/internal/com/logger"
)

var (
	rc *redis.Ring
)

func init() {
	addrs := map[string]string{}
	for _, addr := range getServers() {
		addrs[addr] = addr
	}

	rc = redis.NewRing(&redis.RingOptions{
		Addrs: addrs,
		DB: int64(cfg.C.Rc.Db),
		Password: cfg.C.Rc.Password,
		MaxRetries: 3,
		PoolSize: cfg.C.Rc.PoolSize,
	})
}

func getServers() []string {
	arrServerHosts := make([]string, 0)
	for _, item := range cfg.C.Rc.Servers{
		arr := strings.Split(item, ":")
		if len(arr) < 2 {
			logger.Error("cache.init", logger.Format("err", "Redis servers 配置错误", "item", item))
			continue
		}

		arrServerHosts = append(arrServerHosts, fmt.Sprintf("%s:%s", arr[0], arr[1]))
	}

	return arrServerHosts
}