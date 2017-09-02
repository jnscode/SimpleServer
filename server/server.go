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

func procSession(conn net.Conn) {
	println("session start")

	defer func() {
		println("session end")
		conn.Close()
	}()

	buf := make([]byte, 50)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("connect closed,", err.Error())
			return
		}

		msg := string(buf[0:n])
		fmt.Println("recv msg:", msg)

		conn.Write([]byte(msg + " reply"))
	}
}

func main() {
	fmt.Println("process start")

	addr := "0.0.0.0:7777"

	if len(os.Args) >= 2 {
		addr = os.Args[1]
		fmt.Println("use arg addr ", addr)
	}

	sockListen, err := net.Listen("tcp", addr)
	procError("Create socket error", err)

	defer func() {
		println("close listen sock")
		sockListen.Close()
	}()

	for {
		conn, err := sockListen.Accept()
		procError("Error on accept connection", err)

		println("new connection established")
		go procSession(conn)
	}

	println("process exit")
}
