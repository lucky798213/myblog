package routes

import (
	"github.com/gin-gonic/gin"
	"myblog/utils"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 200,
				"msg": "hello gin",
			})
		})
	}
	r.Run(utils.HttpPort)
}
