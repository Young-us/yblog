package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"

	_ "goblog/config"
	"goblog/pkg/config"
	"net/http"
)

/* func init() {
	// 初始化配置信息
	config.Initialize()
} */

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	//port := os.Args[1]
	http.ListenAndServe(":"+config.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	//http.ListenAndServe(":"+port, middlewares.RemoveTrailingSlash(router))
}
