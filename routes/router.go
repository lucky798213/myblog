package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myblog/api/v1"
	"myblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode) //设置应用程序的运行模式
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口

		//文章模块的路由接口
	}
	r.Run(utils.HttpPort)
}
