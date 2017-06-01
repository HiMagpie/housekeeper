package dao

import (
	"housekeeper/internal/com/cfg"
	"fmt"
	"housekeeper/internal/model/third"
	"housekeeper/internal/model/db"
	"github.com/go-ozzo/ozzo-dbx"
	"time"
)

const (
	MSG_STAT_REC = 1
	MSG_STAT_QUEUED = 2
	MSG_STAT_PUSH_ACK = 3
	MSG_STAT_PUSH_SUCC = 4
	MSG_STAT_OFFLINE = 5
	MSG_STAT_OVERDUE = 6
	MSG_STAT_FAIL = 7
)

type MsgInfo struct {
	Id        int64 `json:"id"`
	Cids      string `json:"cids"`
	MsgId     int64 `json:"msg_id"`
	MsgCtime  int64 `json:"msg_ctime"`
	Ring      bool `json:"ring"`
	Vibrate   bool `json:"vibrate"`
	Cleanable bool `json:"cleanable"`
	Trans     int `json:"trans"`
	Begin     string `json:"begin"`
	End       string `json:"end"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Logo      string `json:"logo"`
	Url       string `json:"url"`
	Status    int `json:"status"`
	Enabled   int `json:"enabled"`
	Ctime     string `json:"ctime"`
	Utime     string `json:"utime"`
}

func (this *MsgInfo) Table(msgId int64) string {
	return cfg.C.Db.GetTable("msg_info", fmt.Sprintf("%d", msgId))
}

func NewMsgInfo() *MsgInfo {
	return new(MsgInfo)
}

func (this *MsgInfo) Create(msg *third.Msg, msgId int64, cids string) (int64, error) {
	res, err := db.GetDb().Insert(this.Table(msgId), dbx.Params{
		"cids": cids,
		"msg_id": msgId,
		"msg_ctime": time.Now().Unix(),
		"ring": msg.MsgConf.Ring,
		"vibrate": msg.MsgConf.Vibrate,
		"cleanable": msg.MsgConf.Cleanable,
		"trans": msg.MsgConf.Trans,
		"begin": msg.MsgConf.Begin,
		"end": msg.MsgConf.End,
		"title": msg.Title,
		"text": msg.Text,
		"logo": msg.Logo,
		"url":msg.Url,
		"status": MSG_STAT_QUEUED,
		"ctime":GetNow(),
	}).Execute()
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (this *MsgInfo) GetByMsgId(msgId int64) (*MsgInfo, error) {
	err := db.GetDb().Select("*").From(this.Table(msgId)).Where(dbx.HashExp{"msg_id":msgId}).One(this)
	return this, err
}
