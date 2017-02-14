package main

import (
	"fmt"
	"net"
	"os"
)

func procError(title string, err error) {
	if err != nil {
		fmt.Println(title, ",", err.Error())
		os.Exit(1)
	}
}

func onSession(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 50)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("connect closed,", err.Error())
			return
		}

		fmt.Println("recv msg:", string(buf[0:n]))
	}
}

func main() {
	fmt.Println("server start")

	if len(os.Args) < 2 {
		fmt.Println("invalid args")
		os.Exit(1)
	}

	sockListen, err := net.Listen("tcp", os.Args[1])
	procError("Create socket error", err)

	defer sockListen.Close()

	for {
		conn, err := sockListen.Accept()
		procError("Error on accept connection", err)

		onSession(conn)
	}
}
