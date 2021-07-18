package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

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
