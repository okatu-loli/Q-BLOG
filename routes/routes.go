package routes

import (
	v1 "Q-BLOG/api/v1"
	"Q-BLOG/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		//user模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口

		//文章模块的路由接口
	}
	log.Fatal(r.Run(utils.HttpPort))
}
