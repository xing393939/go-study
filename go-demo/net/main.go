package main

import (
	"net"
)

func clientRedis() {
	conn, _ := net.Dial("tcp", "127.0.0.1:6379")
	_, _ = conn.Write([]byte("ping\r\n"))
	buffer := make([]byte, 16)
	count, _ := conn.Read(buffer)
	println(string(buffer), count)
}

func clientICMP() {
	ipAddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
	conn, _ := net.ListenIP("ip4:icmp", ipAddr)
	buf := make([]byte, 1024)
	n, addr, _ := conn.ReadFrom(buf)
	println(n, addr)
}

func serverTCP() {
	tcpAddr, _ := net.ResolveTCPAddr("", "127.0.0.1:8080")
	listener, _ := net.ListenTCP("tcp4", tcpAddr)
	defer listener.Close()
	for {
		conn, _ := listener.AcceptTCP()
		go func() {
			buffer := make([]byte, 16)
			_, _ = conn.Read(buffer)
			_, _ = conn.Write(append([]byte{':'}, buffer...))
			_ = conn.Close()
		}()
	}
}
