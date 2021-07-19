package main

import (
	"gin-vue-admin/api/ws"
)

func main() {
	//创建定时器，2秒后，定时器就会向自己的C字节发送一个time.Time类型的元素值
	//timer1 := time.NewTimer(time.Second * 2)
	//t1 := time.Now() //当前时间
	//fmt.Printf("t1: %v\n", t1)
	//
	//t2 := <-timer1.C
	//fmt.Printf("t2: %v\n", t2)

	ws.SendMsg("abcdddd")
}
