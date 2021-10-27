package main

import (
	_ "gin_example/init"
	"gin_example/internal/router"
	"net/http"

	"gin_example/global"

	"github.com/gin-gonic/gin"
)

// @title 用户管理
// @version 1.0
// @description 用户管理: 多角色 多权限
// @termsOfService https://github.com
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := router.NewRouter()
	server := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Log.Info("Start Server Success")
	server.ListenAndServe()
}
