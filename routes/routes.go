package routes

import (
	"Q-BLOG/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	log.Fatal(r.Run(utils.HttpPort))
}
