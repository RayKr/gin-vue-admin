package router

import (
	"gin-vue-admin/api/ws"
	"github.com/gin-gonic/gin"
)

func InitWSRouter(Router *gin.RouterGroup) {
	WSRouter := Router.Group("ws")
	{
		WSRouter.GET("echo", ws.EchoMessage) // echo demo
	}
}
