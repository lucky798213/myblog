package main

import (
	"myblog/model"
	"myblog/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}
