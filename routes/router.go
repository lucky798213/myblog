package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myblog/api/v1"
	"myblog/middleware"
	"myblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode) //设置应用程序的运行模式
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)

		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		//上传文件
		auth.POST("upload", v1.UpLoad)

	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCate)
		router.GET("article", v1.GetArticle)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.GET("article/list/:id", v1.GetCateArt)
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
