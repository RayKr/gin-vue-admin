package initialize

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func EchoMessage(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer ws.Close()

	for {
		//读取ws中的数据
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		// 把消息打印到标准输出
		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))

		//写入ws数据
		if err = ws.WriteMessage(mt, msg); err != nil {
			break
		}
	}
}

func WS() {

	if global.GVA_CONFIG.WS.Start {
		fmt.Printf("fsafadsfdsaf")
		bindAddress := "127.0.0.1:" + global.GVA_CONFIG.WS.Port
		r := gin.Default()
		r.GET("/echo", EchoMessage)
		r.Run(bindAddress)
		fmt.Printf("WebSocket监听启动成功，正在监听[%v]端口\n", global.GVA_CONFIG.WS.Port)
	}
}
