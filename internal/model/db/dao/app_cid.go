package dao

import (
	"housekeeper/internal/com/cfg"
	"housekeeper/internal/model/db"
	"github.com/go-ozzo/ozzo-dbx"
	"fmt"
	"github.com/linkosmos/mapop"
	"housekeeper/internal/com/logger"
)

type AppCid struct {
	Id      int64 `json:"id"`
	AppId   string `json:"app_id"`
	Cid     string `json:"cid"`
	Enabled int `json:"enabeld"`
	Ctime   string `json:"ctime"`
	Utime   string `json:"utime"`
}

func (this *AppCid) Table(cid string) string {
	return cfg.C.Db.GetTable("app_cid", cid)
}

func NewAppCid() *AppCid {
	return new(AppCid)
}

// validate submitted cids
func (this *AppCid) ValidCids(appId string, cids []string) (validCids, invalidCids []string, err error) {
	appCids, err := this.GetAppCidByCids(appId, cids)
	cidMap := make(map[string]bool)
	for _, r := range appCids {
		cidMap[r.Cid] = true
	}
	for _, cid := range cids {
		if _, ok := cidMap[cid]; !ok {
			invalidCids = append(invalidCids, cid)
		}else {
			validCids = append(validCids, cid)
		}
	}

	return validCids, invalidCids, nil
}

// Get cids by app_id and cids, for filtering validated cids
func (this *AppCid) GetAppCidByCids(appId string, cids []string) ([]AppCid, error) {
	tableCids := this.GatherCidToTable(cids)

	// get cid by sql union
	rows := make([]AppCid, 0)
	params := dbx.Params{"app_id": appId}
	selects := make([]string, 0, len(tableCids))
	for table, tCids := range tableCids {
		inStr, tp := makeSqlIn(table, "cid", tCids)
		params = mapop.Merge(params, tp)

		tmpSql := fmt.Sprintf("select cid from %s where app_id = {:app_id} and cid in (%s)", table, inStr)
		selects = append(selects, db.GetDb().NewQuery(tmpSql).SQL())
	}

	err := makeUnionQuery(selects, params).All(&rows)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// gather cid by table name
func (this *AppCid) GatherCidToTable(cids []string) map[string][]string {
	tableCids := make(map[string][]string)
	for _, cid := range cids {
		if _, exist := tableCids[this.Table(cid)]; !exist {
			tableCids[this.Table(cid)] = make([]string, 0)
		}

		tableCids[this.Table(cid)] = append(tableCids[this.Table(cid)], cid)
	}

	return tableCids
}

// check whether record exist while create new one
func (this *AppCid) CreateUnique(cid, appId string) (int64, error) {
	appCid, _ := this.GetByCidAppId(cid, appId)
	if appCid != nil {
		return appCid.Id, nil
	}

	res, _ := db.GetDb().Insert(this.Table(cid), dbx.Params{
		"cid": cid,
		"app_id": appId,
		"ctime": GetNow(),
	}).Execute()
	return res.LastInsertId()
}

func (this *AppCid) GetByCidAppId(cid, appId string) (*AppCid, error) {
	err := db.GetDb().Select("*").From(this.Table(cid)).Where(dbx.HashExp{"cid":cid, "app_id":appId}).One(this)
	if err != nil {
		logger.Error("app_cid", logger.Format("err", err.Error(), "cid", cid, "app_id", appId))
	}
	return this, err
}

