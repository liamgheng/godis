package server

import (
	"context"
	"log"
	"net"

	"github.com/liamgheng/godis/database"
	"github.com/liamgheng/godis/lib/atomic"
	"github.com/liamgheng/godis/redis/parser"
)

type Handler struct {
	closing atomic.Boolean
	db      database.DB
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	doneCh := make(chan struct{})

	payloadCh := parser.ParseStream(conn)
	for payload := range payloadCh {
		if h.closing.Get() {
			return
		}

		go func() {
			// err != nil 时， channel closed
			if payload.Err != nil {
				log.Println(payload.Err)
				return
			}

			result, err := h.db.Exec(ctx, payload.Data)
			if err != nil {
				log.Println(err)
				return
			}
			_, _ = conn.Write(result.ToBytes())
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

func (h *Handler) Close() error {
	h.closing.Set(true)
	return nil
}
