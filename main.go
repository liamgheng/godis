package main

import (
	"fmt"
	"log"
	"net"

	"github.com/liamgheng/godis/redis/server"
	"github.com/liamgheng/godis/tcp"
)

func main() {
	// TODO 需要初始化 db
	handler := &server.Handler{}

	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	addr := listener.Addr().String()
	log.Println(fmt.Sprintf("starting lisen %v", addr))

	tcp.ListenAndServe(listener, handler)
}
