package main

import (
	"Q-BLOG/model"
	"Q-BLOG/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}
