package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func procError(title string, err error) {
	if err != nil {
		fmt.Println(title, ",", err.Error())
		os.Exit(1)
	}
}

func procConnect(wg *sync.WaitGroup, addr string, id int) {
	sock, err := net.Dial("tcp", addr)
	procError("error", err)

	defer func() {
		println("session end")
		sock.Close()
		wg.Done()
	}()

	for i := 0; i < 1000; i++ {
		msg := "hello" + strconv.Itoa(id)
		println("send:", msg)
		sock.Write([]byte(msg))

		buf := make([]byte, 50)
		n, e := sock.Read(buf)
		procError("error", e)

		fmt.Println("recv:", string(buf[0:n]))
		time.Sleep(time.Second)
	}
}

func procInteractive(addr string) {
	sock, err := net.Dial("tcp", addr)
	procError("error", err)

	defer func() {
		println("session end")
		sock.Close()
	}()

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

func main() {
	fmt.Println("process start")

	addr := "0.0.0.0:7777"

	if len(os.Args) >= 2 {
		addr = os.Args[1]
		fmt.Println("use arg addr ", addr)
	}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go procConnect(&wg, addr, i+1)
	}

	wg.Wait()

	println("process exit")
}
