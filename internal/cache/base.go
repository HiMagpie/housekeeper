package cache
import (
	"time"
	"errors"
	"gopkg.in/redis.v3"
	"housekeeper/internal/com/logger"
)

func PushToQueue(queue, item string) (int64, error) {
	c := rc.LPush(queue, item)
	return c.Result()
}

func PopFromQueue(queue string) (string, error) {
	c := rc.RPop(queue)
	return c.Result()
}

func BPopFromQueue(queue string, timeout time.Duration) (string, error) {
	c := rc.BRPop(timeout, queue)
	res, err := c.Result()
	if err != nil {
		if err != redis.Nil {
			logger.Error("redis.error", err.Error())
		}

		return "", err
	}

	if len(res) < 2 {
		return "", errors.New("Bpop from queue encounter invalid value.")
	}

	return res[1], nil
}

func BRPopFromQueue(queue string, timeout time.Duration) (string, error) {
	return BPopFromQueue(queue, timeout)
}

func HSet(key, field, val string) error {
	c := rc.HSet(key, field, val)
	return c.Err()
}

func SAdd(key string, v ...string) error {
	c := rc.SAdd(key, v...)
	return c.Err()
}