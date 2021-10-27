package init

import (
	"fmt"
	"gin_example/common/setting"
	"gin_example/global"
	"log"
	"time"
)

func init() {
	err := initSetting()
	if err != nil {
		log.Fatalf("init.initSetting err: %v", err)
	}
	initLog()
	err = initMysql()
	if err != nil {
		global.Log.Error(fmt.Sprintf("init.initMysql %v", err))
	}
}

func initSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Mysql", &global.MysqlSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Log", &global.LogSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
