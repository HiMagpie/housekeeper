package main

import (
	"housekeeper/internal/com/logger"
	"housekeeper/internal/com/cfg"
	_ "housekeeper/internal/model"
	"housekeeper/internal/com/drivers/db/orm"
	"fmt"
	"os"
)

func main() {
	logger.Info("init.table", cfg.C.Db.SliceTable)

	initTables()
}

func initTables() {
	initApp(getSliceTableNum("app"))
	initAppCid(getSliceTableNum("app_cid"))
	initMsgCidStatus(getSliceTableNum("msg_cid_status"))
	initMsgInfo(getSliceTableNum("msg_info"))
}

func initApp(num int) {
	createSql := `CREATE TABLE %sapp%s (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(40) NOT NULL COMMENT '应用名称',
    app_id varchar(32) NOT NULL DEFAULT '' COMMENT 'app唯一标示',
    app_secret varchar(32) NOT NULL DEFAULT '' COMMENT 'app密钥，用于手机客户端',
    master_secret varchar(32) NOT NULL DEFAULT '' COMMENT 'app密钥，用于服务端向HiMagpie发消息密钥',
    enabled tinyint(1) NOT NULL DEFAULT '1',
    ctime timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
    utime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY app_id (app_id,app_secret,master_secret)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	dropSql := "DROP TABLE IF EXISTS %sapp%s;"

	initTableWithTmpSql(dropSql, createSql, num)
}

func initAppCid(num int) {
	createSql := `CREATE TABLE %sapp_cid%s (
	id int(11) unsigned NOT NULL AUTO_INCREMENT,
	app_id varchar(32) NOT NULL DEFAULT '' COMMENT '第三方应用唯一表示',
	cid varchar(32) NOT NULL DEFAULT '' COMMENT 'cid',
	enabled tinyint(1) NOT NULL DEFAULT '1',
	ctime timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
	utime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第三方应用和 cid 的绑定关系';`
	dropSql := "DROP TABLE IF EXISTS %sapp_cid%s;"

	initTableWithTmpSql(dropSql, createSql, num)
}

func initMsgCidStatus(num int) {
	createSql := `CREATE TABLE %smsg_cid_status%s (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    cid varchar(32) NOT NULL DEFAULT '' COMMENT '对应客户端id,',
    msg_id bigint(20) NOT NULL DEFAULT '0' COMMENT '该消息是该cid对应的msg_id（作为分表依赖的字段）',
    status tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '1：收到；2：放到目标服务器队列；3：服务器已推送，等待ACK；4：收到ACK，成功；5：客户端不在线，消息离线中；6：消息超时，失效；7：失败；',
    enabled tinyint(4) NOT NULL DEFAULT '1',
    ctime timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
    utime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY cid_msg_index (cid,msg_id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	dropSql := "DROP TABLE IF EXISTS %smsg_cid_status%s;"

	initTableWithTmpSql(dropSql, createSql, num)
}

func initMsgInfo(num int) {
	createSql := `CREATE TABLE %smsg_info%s (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    msg_id bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '给该消息分配的id（作为分表依赖的字段）',
    cids text NOT NULL COMMENT 'cid客户端id,多个用,分隔',
    msg_ctime int(11) NOT NULL COMMENT '收到该消息的时间',
    ring tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '消息铃声',
    vibrate tinyint(4) unsigned NOT NULL DEFAULT '0',
    cleanable tinyint(4) unsigned NOT NULL DEFAULT '1',
    trans tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '传输方式1：透传；2： 提醒',
    begin varchar(8) NOT NULL DEFAULT '00:00:00' COMMENT '接收的开始时间',
    end varchar(8) NOT NULL DEFAULT '00:00:00' COMMENT '接受的结束时间',
    title varchar(500) NOT NULL DEFAULT '' COMMENT '标题栏信息',
    text varchar(1000) NOT NULL DEFAULT '' COMMENT '正文',
    logo varchar(200) NOT NULL DEFAULT '' COMMENT '通知提醒的icon地址',
    url varchar(200) NOT NULL DEFAULT '' COMMENT '跳转网页的url',
    status tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '1：收到；2：全部发送成功；3：部分失败；4：完全失败；',
    enabled tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '0: 删除; 1: 正常',
    ctime timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
    utime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	dropSql := "DROP TABLE IF EXISTS %smsg_info%s;"
	initTableWithTmpSql(dropSql, createSql, num)
}

func initTableWithTmpSql(dropSql, createSql string, num int) {
	for i := 0; i < num; i++ {
		suffix := getSuffix(i, num)
		dropSql := fmt.Sprintf(dropSql, cfg.C.Db.Prefix, suffix)
		err := querySql(dropSql)
		if err != nil {
			os.Exit(-1)
		}

		tmpSql := fmt.Sprintf(createSql, cfg.C.Db.Prefix, suffix)
		err = querySql(tmpSql)
		if err != nil {
			os.Exit(-1)
		}
	}
}

func getSliceTableNum(tableName string) int {
	num := 1
	if _, ok := cfg.C.Db.SliceTable[tableName]; ok {
		num = cfg.C.Db.SliceTable[tableName]
	}

	return num
}

func querySql(sqlStr string) error {
	q, err := orm.GetDB(cfg.C.Db.Alias)
	if err != nil {
		logger.Error("init.table", logger.Format("err", err.Error()))
		panic(err)
	}
	_, err = q.Exec(sqlStr)
	if err != nil {
		logger.Error("init.table", logger.Format("err", err.Error()), sqlStr)
		return err
	}

	logger.Info("init.table", logger.Format("sql", sqlStr))
	return nil
}

func getSuffix(i, num int) string {
	suffix := ""
	if num > 1 {
		suffix = fmt.Sprintf("_%d", i)
	}

	return suffix
}
