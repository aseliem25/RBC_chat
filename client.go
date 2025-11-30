package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run client.go <server:port>")
		return
	}

	conn, _ := net.Dial("tcp", os.Args[1])

	// الاستقبال
	go func() {
		sc := bufio.NewScanner(conn)
		for sc.Scan() {
			fmt.Println(sc.Text())
		}
		os.Exit(0)
	}()

	// الإرسال
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fmt.Fprintln(conn, sc.Text())
	}
}
