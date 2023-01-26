package main

import (
	"fmt"
	"log"
	"net"

	"github.com/liamgheng/godis/tcp"
)

func main() {
	handler := &tcp.EchoHandler{}

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	addr := listener.Addr().String()
	log.Println(fmt.Sprintf("start lisen %v", addr))

	tcp.ListenAndServe(listener, handler)
}
