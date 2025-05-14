package receiver

import (
	"bufio"
	"net"

	"github.com/go-redis/redis/v8"
)

var LogChannel = make(chan string, 100)

func StartLogReceiver(port string, rdb *redis.Client) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			go handleConnection(conn)
		}
	}()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		LogChannel <- line
	}
}
