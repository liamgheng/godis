// 用来测试 tcp server 是否正常工作

package tcp

import (
	"bufio"
	"context"
	"log"
	"net"
	"time"

	"github.com/liamgheng/godis/lib/atomic"
)

type EchoHandler struct {
	closing atomic.Boolean
}

func (h *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	doneCh := make(chan struct{})

	for {
		if h.closing.Get() {
			return
		}
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		go func() {
			b := []byte(msg)
			_, _ = conn.Write(b)
			doneCh <- struct{}{}
		}()

		select {
		case <-doneCh:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func (h *EchoHandler) Close() error {
	h.closing.Set(true)
	return nil
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
