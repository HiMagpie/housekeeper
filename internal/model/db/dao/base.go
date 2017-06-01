package dao

import (
	"github.com/go-ozzo/ozzo-dbx"
	"fmt"
	"strings"
	"housekeeper/internal/model/db"
	"housekeeper/internal/com/utils"
)

const (
	ENABLED = 1
	DELETED = 0
)

//return yyyy-mm-dd hh:ii:ss
func GetNow() string {
	return utils.CurDateTime()
}

// key - unique key
// field - sql field name
// data - bind data
func makeSqlIn(key, field string, data []string) (string, dbx.Params) {
	inStr := ""
	p := dbx.Params{}
	for i, v := range data {
		tk := fmt.Sprintf("%s_%s_%d", key, field, i)
		inStr += fmt.Sprintf("{:%s},", tk)
		p[tk] = v
	}

	return strings.TrimRight(inStr, ","), p
}

// selects - select sql
// p - bind values
func makeUnionQuery(selects []string, p dbx.Params) *dbx.Query {
	sqlStr := ""
	for _, s := range selects {

		sqlStr += fmt.Sprintf("(%s) UNION ", s)
	}
	sqlStr = strings.TrimRight(sqlStr, "UNION ")

	q := db.GetDb().NewQuery(sqlStr).Bind(p)
	q.LogFunc = db.GetDb().LogFunc
	return q
}
