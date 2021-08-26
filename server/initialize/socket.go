package initialize

import (
	"bufio"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"net"
)

func process(conn net.Conn) {
	// 处理完关闭连接
	defer conn.Close()

	fmt.Printf("有新设备连接成功：%v\n", conn.RemoteAddr())

	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}

		rev := string(buf[:n])
		fmt.Printf("收到[%v]的数据：%v\n", conn.RemoteAddr(), rev)

		// 将接受到的数据返回给客户端
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}

func Socket() {
	if global.GVA_CONFIG.Socket.Start {
		// 建立 tcp 服务
		listen, err := net.Listen("tcp", global.GVA_CONFIG.Socket.Addr)
		if err != nil {
			fmt.Printf("listen failed, err:%v\n", err)
			return
		}

		fmt.Printf("Socket监听启动成功，正在监听：%v\n", global.GVA_CONFIG.Socket.Addr)

		for {
			// 等待客户端建立连接
			conn, err := listen.Accept()
			if err != nil {
				fmt.Printf("accept failed, err:%v\n", err)
				continue
			}
			// 启动一个单独的 goroutine 去处理连接
			go process(conn)
		}
	}
}
