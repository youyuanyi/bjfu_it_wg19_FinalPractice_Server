package routes

import (
	"WeatherServer/controller"
	"WeatherServer/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func UserRoutes(e *gin.Engine) *gin.Engine {
	//用户注册
	e.POST("/register", controller.Register)
	//用户登录
	e.POST("/login", controller.Login)
	//登录获取用户信息
	return e
}

func UserMgrRoutes(e *gin.Engine) *gin.Engine {
	userRoutes := e.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware())
	userRoutes.GET("", controller.GetInfo)           // 验证用户
	userRoutes.POST("users", controller.GetAllUsers) // 获取所有用户信息
	userRoutes.POST("", controller.AddUser)          // 添加用户
	userRoutes.PUT(":id", controller.EditUser)       // 编辑用户
	userRoutes.DELETE(":id", controller.DelUser)     // 删除用户
	return e
}

func NodeRoutes(e *gin.Engine) *gin.Engine {
	nodeRoutes := e.Group("/node")
	nodeRoutes.Use(middleware.AuthMiddleware())
	nodeRoutes.GET("", controller.GetAllNode)            // 获取所有设备
	nodeRoutes.GET(":id", controller.ListNodes)          // 展示当前用户的设备
	nodeRoutes.POST("", controller.AddNode)              // 添加用户
	nodeRoutes.DELETE(":id", controller.DeleteNode)      // 删除设备,id是设备id
	nodeRoutes.PUT(":id", controller.EditNode)           // 修改设备,id是设备id
	nodeRoutes.POST("setTime", controller.SetSystemTime) // 设置系统时间
	return e
}

func PhyRoutes(e *gin.Engine) *gin.Engine {
	phyRoutes := e.Group("/phy")
	phyRoutes.Use(middleware.AuthMiddleware())
	phyRoutes.POST(":id", controller.GetPhy)              // 获取所有的物理量
	phyRoutes.PUT(":id", controller.EditPhy)              // 修改物理量
	phyRoutes.GET("/node/:id", controller.GetPhyByNodeID) // 根据NodeID获取对应的物理量名
	return e
}
func DataRoutes(e *gin.Engine) *gin.Engine {
	dataRoutes := e.Group("/data")
	dataRoutes.Use(middleware.AuthMiddleware())
	dataRoutes.POST(":id", controller.ShowData)
	dataRoutes.GET("/download/:id", controller.DownloadExcel)
	return e
}
func CollectRoutes(e *gin.Engine) *gin.Engine {
	//允许跨域访问
	e.Use(middleware.CORSMiddleware())
	e = UserRoutes(e)
	e = UserMgrRoutes(e)
	e = DataRoutes(e)
	e = NodeRoutes(e)
	e = PhyRoutes(e)
	e.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	return e
}
