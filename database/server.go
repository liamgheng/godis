package database

import (
	"context"

	"github.com/liamgheng/godis/redis/protocol"
)

// Resp 是格式化的数据，需要被转为 raw response
// 每一种类型不一样，所以用接口来定义
type Resp interface {
	ToBytes() []byte
}

type DB struct {
}

func (s *DB) Exec(ctx context.Context, req protocol.MultiBulk) (Resp, error) {
	return nil, nil
}
