package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/ws"
	"github.com/gin-gonic/gin"
)

func InitWSRouter(Router *gin.RouterGroup) {
	WSRouter := Router.Group("ws")
	{
		// echo demo
		WSRouter.GET("echo", ws.EchoMessage)
		// 广播功能
		WSRouter.GET("broadcast", ws.Broadcast)
		// 测试
		WSRouter.GET("send", ws.Send)
	}
}
