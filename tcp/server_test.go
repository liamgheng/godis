package tcp

import (
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenAndServe(t *testing.T) {
	handler := &EchoHandler{}

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	addr := listener.Addr().String()
	log.Println(fmt.Sprintf("start lisen %v", addr))

	go ListenAndServe(listener, handler)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	client := EchoClient{conn: conn}
	res := client.Ping()

	assert.Equal(t, "PING\n", res)
}
