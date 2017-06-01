package dao

import (
	"housekeeper/internal/com/cfg"
	"housekeeper/internal/model/db"
	"github.com/go-ozzo/ozzo-dbx"
)

type MsgCidStatus struct {
	Id      int64 `json:"id"`
	MsgId   int64 `json:"msg_id"`
	Cid     string `json:"cid"`
	Status  int `json:"status"`
	Enabled int `json:"enabled"`
	Ctime   string `json:"ctime"`
	Utime   string `json:"utime"`
}

func (this *MsgCidStatus) Table(cid string) string {
	return cfg.C.Db.GetTable("msg_cid_status", cid)
}

func NewMsgCidStatus() *MsgCidStatus {
	return new(MsgCidStatus)
}

// create new record
func (this *MsgCidStatus) Create(cid string, msgId int64) (int64, error) {
	res, _ := db.GetDb().Insert(this.Table(cid), dbx.Params{
		"msg_id": msgId,
		"cid": cid,
		"status": MSG_STAT_QUEUED,
		"ctime": GetNow(),
	}).Execute()
	return res.LastInsertId()
}

// update cid <-> msg_id's status
func (this *MsgCidStatus) UpdateStatusByCidMsgId(cid string, msgId int64, status int) error {
	_, err := db.GetDb().Update(this.Table(cid), dbx.Params{"status": status}, dbx.HashExp{
		"cid":cid,
		"msg_id":msgId,
	}).Execute()
	return err
}
