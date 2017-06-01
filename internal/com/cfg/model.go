package cfg

import (
	"github.com/larspensjo/config"
	"log"
	"strings"
	"strconv"
	"housekeeper/internal/com/utils"
	"fmt"
)

type Conf struct {
	HttpSrv *HttpSrv
	RpcSrv  *RpcSrv
	Rc      *Rc
	Db      *Db
	Log     *Log
}

func NewConf() *Conf {
	return &Conf{
		HttpSrv: &HttpSrv{},
		RpcSrv:&RpcSrv{},
		Rc: &Rc{},
		Db:&Db{},
		Log: &Log{},
	}
}

type HttpSrv struct {
	Port              int
	SnowFlakeWorkerId int
}

type RpcSrv struct {
	Port         int
	PushServerHb int
}

type Rc struct {
	Servers     []string
	Password    string
	Db          int
	PoolSize    int
	PoolTimeout int
	VnodeNum    int
}

type Db struct {
	Alias      string
	Driver     string
	Host       string
	Port       int
	Database   string
	Username   string
	Password   string
	Charset    string
	Prefix     string
	Debug      bool

	SliceTable map[string]int
}

type Log struct {
	ChannelLen int
	Path       string
	Level      int
}

func parseConfigFile(path string) {
	cfg, err := config.ReadDefault(path)
	if err != nil {
		log.Fatalln("[E] ", err.Error(), " cfg file: ", path)
	}

	parseHttpSrv(cfg)
	parseRpcSrv(cfg)
	parseRc(cfg)
	parseDb(cfg)
	parseLog(cfg)
}

func validSection(cfg *config.Config, section string) {
	if !cfg.HasSection(section) {
		log.Fatalf("[E] config options of %s not found!\n", section)
	}
}

func parseHttpSrv(cfg *config.Config) *HttpSrv {
	section := "server"
	validSection(cfg, section)

	c := C.HttpSrv
	c.Port, _ = cfg.Int(section, "port")
	// @TODO record worker id and check whether exist
	// @TODO cron update worker_id list and remove invalid worker id
	// @TODO check worker_id range (0 < worker_id < 1024)
	c.SnowFlakeWorkerId, _ = cfg.Int(section, "snowflake_worker_id")
	return c
}

func parseRpcSrv(cfg *config.Config) *RpcSrv {
	section := "rpc"
	validSection(cfg, section)

	c := C.RpcSrv
	c.Port, _ = cfg.Int(section, "port")
	c.PushServerHb, _ = cfg.Int(section, "push_server_hb")
	return c
}

func parseRc(cfg *config.Config) *Rc {
	section := "rc"
	validSection(cfg, section)

	c := C.Rc
	servers, _ := cfg.String(section, "servers")
	c.Servers = strings.Split(strings.Trim(servers, "|"), "|")
	c.Password, _ = cfg.String(section, "password")
	c.Db, _ = cfg.Int(section, "db")
	c.PoolSize, _ = cfg.Int(section, "pool_size")
	c.PoolTimeout, _ = cfg.Int(section, "pool_timeout")
	c.VnodeNum, _ = cfg.Int(section, "vnod_num")
	return c
}

func parseDb(cfg *config.Config) *Db {
	section := "db"
	validSection(cfg, section)

	c := C.Db
	c.Alias, _ = cfg.String(section, "alias")
	c.Driver, _ = cfg.String(section, "driver")
	c.Host, _ = cfg.String(section, "host")
	c.Port, _ = cfg.Int(section, "port")
	c.Database, _ = cfg.String(section, "database")
	c.Username, _ = cfg.String(section, "username")
	c.Password, _ = cfg.String(section, "password")
	c.Charset, _ = cfg.String(section, "charset")
	c.Prefix, _ = cfg.String(section, "prefix")
	c.Debug, _ = cfg.Bool(section, "debug")

	c.SliceTable = make(map[string]int, 0)
	st, _ := cfg.String(section, "slice_table")
	st = strings.Trim(st, ",")
	if st != "" {
		tbConfs := strings.Split(st, ",")
		for _, tbc := range tbConfs {
			ret := strings.Split(tbc, ":")
			if len(ret) != 2 {
				panic("[e] db config \"slice_table\" format invalid: " + st)
			}

			// table_name => slice table num
			tableNum, err := strconv.Atoi(ret[1])
			if err != nil {
				panic("[e] db config \"slice_table\" format invalid: " + err.Error())
			}
			c.SliceTable[ret[0]] = tableNum
		}
	}

	return c
}

// Get real table name by slice field value.
func (this *Db) GetTable(rawName, key string) string {
	if _, ok := this.SliceTable[rawName]; !ok {
		return this.Prefix + rawName
	}

	appIdCrc, _ := utils.Crc32Str(key)
	suffix := appIdCrc % uint32(this.SliceTable[rawName])
	return fmt.Sprintf(this.Prefix + rawName + "_%d", suffix)
}

func parseLog(cfg *config.Config) *Log {
	section := "log"
	validSection(cfg, section)

	c := C.Log
	c.ChannelLen, _ = cfg.Int(section, "channel_len")
	c.Path, _ = cfg.String(section, "path")
	c.Level, _ = cfg.Int(section, "level")

	return c
}
