package cache

const (
	KEY_MAP_CID_SERVER_NOTIFY_QUEUE = "hmap_cid_server_notify_queue"
)

func GetCidPushServerQueueKey(cid string) string {
	return KEY_MAP_CID_SERVER_NOTIFY_QUEUE + "_" + cid[0:3]
}

func GetServerNotifyQueueByCid(cid string) (string, error) {
	c := rc.HGet(GetCidPushServerQueueKey(cid), cid)
	return c.Result()
}
