package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients = make(map[*Client]bool) // 用户组映射
	join    = make(chan *Client, 10) // 用户加入通道
	leave   = make(chan *Client, 10) // 用户退出通道
	message = make(chan Message, 10) // 消息通道
)

type Message struct {
	EventType byte   `json:"type"`    // 0表示用户发布消息；1表示用户进入；2表示用户退出
	Name      string `json:"name"`    // 用户名称
	Message   string `json:"message"` // 消息内容
}

func init() {
	go broadcaster()
}

//处理每个连接，每个连接，go都会重新起一个协程来处理
func Broadcast(c *gin.Context) {
	w := c.Writer
	r := c.Request
	var (
		wsConn *websocket.Conn
		err    error
		client *Client
		data   []byte
	)

	//返回一个map,并且赋值给r.Form
	e := r.ParseForm()
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(r.Form, r.RemoteAddr)
	var (
		name string
		id   string
	)
	if r.Form == nil || len(r.Form) == 0 {
		name = r.RemoteAddr
		id = r.RemoteAddr
	} else {
		name = r.Form["name"][0]
		id = r.Form["id"][0]
	}

	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if client, err = InitConnection(wsConn); err != nil {
		goto ERR
	}
	client.Id = id
	client.Name = name

	// 如果用户列表中没有该用户
	if !clients[client] {
		join <- client
	}

	for {
		if data, err = client.ReadMessage(); err != nil { //一直读消息，没有消息就阻塞
			goto ERR
		}
		var msg Message
		msg.EventType = 0
		msg.Name = client.Name
		msg.Message = string(data)
		message <- msg
	}

ERR:
	leave <- client //这个客户断开
	client.Close()

}

func broadcaster() {
	for {
		select {
		// 消息通道中有消息则执行，否则堵塞
		case msg := <-message:
			// 将数据编码成json形式，data是[]byte类型
			// json.Marshal()只会编码结构体中公开的属性(即大写字母开头的属性)
			data, err := json.Marshal(msg)
			if err != nil {
				return
			}
			for client := range clients {
				if client.IsClosed == true {
					leave <- client //这个客户断开
					continue
				}
				// fmt.Println("=======the json message is", string(data))  // 转换成字符串类型便于查看
				if client.WriteMessage(data) != nil {
					continue //发送失败就跳过
				}
			}

		// 有用户加入
		case client := <-join:
			clients[client] = true // 将用户加入映射
			// 将用户加入消息放入消息通道
			var msg Message
			msg.Name = client.Name
			msg.EventType = 1
			msg.Message = fmt.Sprintf("%s join in, there are %d preson in room", client.Name, len(clients))
			message <- msg

		// 有用户退出
		case client := <-leave:
			// 如果该用户已经被删除
			if !clients[client] {
				break
			}
			delete(clients, client) // 将用户从映射中删除
			// 将用户退出消息放入消息通道
			var msg Message
			msg.Name = client.Name
			msg.EventType = 2
			msg.Message = fmt.Sprintf("%s leave, there are %d preson in room", client.Name, len(clients))
			message <- msg
		}
	}
}

// 后台发送广播数据
func SendMsg(str string) {
	var msg Message
	msg.EventType = 0
	msg.Name = "ADMIN"
	msg.Message = str
	message <- msg
}
