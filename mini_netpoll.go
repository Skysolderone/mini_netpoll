package main

import "net"

func main() {
	l, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := l.Accept()
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	var buf []byte
	_, _ = conn.Read(buf)
	_, _ = conn.Write(buf)
}
