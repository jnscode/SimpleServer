package main

import (
	"bufio"
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

func main() {
	fmt.Println("client start")

	if len(os.Args) < 2 {
		fmt.Println("invalid args")
		os.Exit(1)
	}

	sock, err := net.Dial("tcp", os.Args[1])
	procError("error", err)

	defer sock.Close()

	sock.Write([]byte("hello"))

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, _ := reader.ReadLine()
		strLine := string(line)

		if strLine == "exit" {
			break
		}

		sock.Write(line)
	}

}
