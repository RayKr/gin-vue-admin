package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "net/http"
)

func BroadCast(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer ws.Close()

	//写入ws数据
	if err = ws.WriteMessage(websocket.TextMessage, []byte("refresh")); err != nil {
		return
	}
}
