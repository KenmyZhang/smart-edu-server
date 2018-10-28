package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"smart-edu-server/app"
	"smart-edu-server/common/util"
	"smart-edu-server/log"
)

func main() {
	log.Info(fmt.Sprintf("Current version is %v (%v/%v)", util.CurrentVersion, util.BuildDate, util.BuildHash))
	server := app.NewServer()
	server.Start()
	server.ShutDown()
	log.Info("shutdown success")
}
