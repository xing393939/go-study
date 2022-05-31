package main

import (
	"net"
)

func testRedis() {
	conn, _ := net.Dial("tcp", "127.0.0.1:6379")
	_, _ = conn.Write([]byte("ping\r\n"))
	buffer := make([]byte, 16)
	count, _ := conn.Read(buffer)
	println(string(buffer), count)
}

func main() {
	ipAddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
	conn, _ := net.ListenIP("ip4:icmp", ipAddr)
	buf := make([]byte, 1024)
	n, addr, _ := conn.ReadFrom(buf)
	println(n, addr)
}
