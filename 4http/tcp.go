package main

import "net"

func tcpHandler(conn net.Conn) {
	_, err := conn.Write([]byte("hello world"))
	if err != nil {
		return
	}

	defer conn.Close()
}

func main() {

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		go tcpHandler(conn)
	}

}
