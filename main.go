package main

import (
	"WeatherServer/common"
	_ "WeatherServer/docs" // swag init 生成的doc路径
	"WeatherServer/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

// @title 天气管理系统后端API文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:9027
// @BasePath /
func main() {
	// 获取初始化的数据库
	db := common.InitDB()
	// 延迟关闭数据库
	defer db.Close()
	// 创建路由引擎
	r := gin.Default()
	// 配置静态文件路径
	r.Static("/assets", "./assets")
	r.StaticFS("/images", http.Dir("./static/images"))
	// 启动路由
	routes.CollectRoutes(r)
	// 启动服务
	panic(r.Run(":9027"))
}
