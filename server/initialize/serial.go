package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tarm/serial"
	"log"
	"time"
)

func Serial() {
	fmt.Println(global.GVA_CONFIG.Serial.Com)
	c := &serial.Config{Name: global.GVA_CONFIG.Serial.Com, Baud: 115200, ReadTimeout: time.Minute * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("【%v】串口打开成功，正在监听...\n", global.GVA_CONFIG.Serial.Com)

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 128)
		n, err = s.Read(buf)
		// TODO 串口数据的处理逻辑
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])
	}
}
