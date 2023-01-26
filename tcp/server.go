// tcp server
package tcp

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}

func ListenAndServe(listener net.Listener, handler Handler) {
	acceptErrCh := make(chan struct{}, 1)
	closeCh := make(chan os.Signal, 1)
	signal.Notify(closeCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case err := <-acceptErrCh:
			log.Println(fmt.Sprintf("accept error: %v", err))
		case <-closeCh:
			log.Println("got exit signal")
		}
		go func() {
			time.AfterFunc(5*time.Second, cancel)
		}()

		log.Println("shutting down...")
		_ = listener.Close()
		_ = handler.Close()
	}()

	// wg 确保所有 G 都执行结束之后才退出
	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			acceptErrCh <- struct{}{}
			break
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			handler.Handle(ctx, conn)
		}()
	}
	wg.Wait()
}
