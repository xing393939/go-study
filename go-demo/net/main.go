package main

import (
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:6379")
	_, _ = conn.Write([]byte("ping\r\n"))
	buffer := make([]byte, 16)
	count, _ := conn.Read(buffer)
	println(string(buffer), count)
}
