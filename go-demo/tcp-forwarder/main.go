package main

import (
	"fmt"
	"net"
	"time"
)

func serve(conn1 net.Conn, conn2 net.Conn) {
	buf := make([]byte, 10240, 10240)
	for {
		n, err := conn1.Read(buf)
		if err != nil {
			_ = conn1.Close()
			fmt.Println("read", err.Error())
			break
		}
		if n > 0 {
			time.Sleep(time.Second * 1) // sleep 10秒
			_, err = conn2.Write(buf[0:n])
			if err != nil {
				_ = conn2.Close()
				fmt.Println("write", err.Error())
				break
			}
		}
	}
}

func main() {
	listener, _ := net.Listen("tcp", "localhost:3307")
	for {
		time.Sleep(time.Second * 10) // sleep 10秒再才连接上
		upstream, _ := listener.Accept()
		downstream, err := net.Dial("tcp", "172.27.178.39:3306")
		if err != nil {
			panic("mysql error")
		}
		go serve(upstream, downstream)
		go serve(downstream, upstream)
	}
}
