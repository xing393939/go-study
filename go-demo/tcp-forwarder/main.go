package main

import (
	"fmt"
	"io"
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
	listener, _ := net.Listen("tcp", "172.30.28.130:3307")
	for {
		upstream, _ := listener.Accept() // accept之前三次握手已经完成
		downstream, err := net.Dial("tcp", "172.30.28.130:3306")
		if err != nil {
			println(err.Error())
			continue
		} else {
			println(downstream.LocalAddr().String() + " -> " + downstream.RemoteAddr().String())
		}
		go io.Copy(upstream, downstream)
		go io.Copy(downstream, upstream)
	}
}
