// tcp server
package tcp

import (
	"net"
)

type Handler interface {
	Handle(conn net.Conn)
}

func ListenAndServe(listener net.Listener, handler Handler) error {
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		// TODO: 限制 GoRoutine 数量
		go func() {
			handler.Handle(conn)
		}()
	}
}
