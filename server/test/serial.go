package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "COM3", Baud: 115200, ReadTimeout: time.Minute * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("COM3串口打开成功，正在监听...")

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 128)
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])
	}
}
