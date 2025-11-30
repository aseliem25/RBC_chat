package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

type Client struct {
	id   int64
	conn net.Conn
}

var (
	clients     = make(map[int64]Client)
	clientsLock sync.Mutex
	nextID      int64
	broadcast   = make(chan string)
)

func main() {
	ln, _ := net.Listen("tcp", ":9000")
	fmt.Println("Server running on :9000")

	go broadcaster()

	for {
		conn, _ := ln.Accept()
		id := atomic.AddInt64(&nextID, 1)

		client := Client{id: id, conn: conn}

		clientsLock.Lock()
		clients[id] = client
		clientsLock.Unlock()

		broadcast <- fmt.Sprintf("User %d joined", id)

		go handleClient(client)
	}
}

func broadcaster() {
	for msg := range broadcast {
		clientsLock.Lock()
		for _, client := range clients {
			fmt.Fprintln(client.conn, msg)
		}
		clientsLock.Unlock()
	}
}

func handleClient(c Client) {
	sc := bufio.NewScanner(c.conn)

	for sc.Scan() {
		text := sc.Text()

		// لا نرسل للمرسل
		clientsLock.Lock()
		for id, client := range clients {
			if id != c.id {
				fmt.Fprintf(client.conn, "[%d] %s\n", c.id, text)
			}
		}
		clientsLock.Unlock()
	}

	// عند خروج العميل، نحذفه
	clientsLock.Lock()
	delete(clients, c.id)
	clientsLock.Unlock()

	broadcast <- fmt.Sprintf("User %d left", c.id)
	c.conn.Close()
}
