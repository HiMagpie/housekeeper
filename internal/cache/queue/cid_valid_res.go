package queue

import (
	"encoding/json"
	"housekeeper/internal/cache"
)

const (
	KEY_QUEUE_CID_VALID_RES_PREFIX = "queue_cid_valid_res_"
)

type CidValidRes struct {
	Cid string `json:"cid"`
	Valid int32 `json:"valid"`
}

func NewCidValidRes() *CidValidRes {
	return new(CidValidRes)
}

func (this CidValidRes) getQueue(hostname string) string {
	return KEY_QUEUE_CID_VALID_RES_PREFIX + hostname
}

func (this *CidValidRes) GetCidValidResJson() (string, error) {
	cidValidResJson, err := json.Marshal(this)
	if err != nil {
		return "", err
	}

	return string(cidValidResJson), nil
}

func (this *CidValidRes)PushToQueue(hostname string) error {
	res, err := this.GetCidValidResJson()
	if err != nil {
		return err
	}

	_, err = cache.PushToQueue(this.getQueue(hostname), res)
	return err
}