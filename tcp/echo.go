// 用来测试 tcp server 是否正常工作

package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type EchoHandler struct {
}

func (h EchoHandler) Handle(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if msg == "PING\n" {
			_, err = conn.Write([]byte("PONG\n"))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(fmt.Sprintf("not support %v", msg))
		}
	}
}

type EchoClient struct {
	conn net.Conn
}

func (c EchoClient) Ping() string {
	_, err := c.conn.Write([]byte("PING\n"))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(10 * time.Millisecond)
	reader := bufio.NewReader(c.conn)
	res, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return res
}
