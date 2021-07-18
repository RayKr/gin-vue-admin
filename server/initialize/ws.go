package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/router"
	"github.com/gin-gonic/gin"
)

func WS() {

	if global.GVA_CONFIG.WS.Start {
		r := gin.Default()
		// 初始化ws相关路由
		PublicGroup := r.Group("")
		{
			router.InitWSRouter(PublicGroup)
		}
		if err := r.Run(":" + global.GVA_CONFIG.WS.Port); err != nil {
			return
		}
	}
}
