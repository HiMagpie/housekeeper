package dao

import (
	"housekeeper/internal/com/cfg"
	"housekeeper/internal/model/db"
	"housekeeper/internal/com/logger"
	"github.com/go-ozzo/ozzo-dbx"
)

type App struct {
	Id           int64 `json:"id"`
	Name         string `json:"name"`
	AppId        string `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	MasterSecret string `json:"master_secret"`
	Enabled      int `json:"enabeld"`
	Ctime        string `json:"ctime"`
	Utime        string `json:"utime"`
}

func (this *App) Table(appId string) string {
	return cfg.C.Db.GetTable("app", appId)
}

func NewApp() *App {
	return new(App)
}

func (this *App) ValidApp(appId, appSecret string) (bool, error) {
	app, err := this.GetByParams(appId, dbx.HashExp{"app_id": appId, "app_secret":appSecret})
	if err != nil {
		return false, err
	}

	return app.Id > 0, nil
}

func (this *App) ValidMaster(appId, masterSecret string) (bool, error) {
	app, err := this.GetByParams(appId, dbx.HashExp{"app_id": appId, "master_secret":masterSecret})
	if err != nil {
		return false, err
	}

	return app.Id > 0, nil
}

func (this *App) GetByParams(appId string, params dbx.HashExp) (*App, error) {
	app := NewApp()
	err := db.GetDb().Select("*").From(this.Table(appId)).
	Where(params).
	One(app)
	if err != nil {
		logger.Error("app", logger.Format("err", err.Error()))
		return nil, err
	}

	return app, nil
}
