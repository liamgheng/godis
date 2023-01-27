package parser

import (
	"bufio"
	"bytes"
	"io"
	"strconv"

	"github.com/liamgheng/godis/redis/protocol"
)

// type Reply interface {
// 	ToBytes() []byte
// }

type Payload struct {
	Data protocol.MultiBulk
	Err  error
}

func ParseStream(r io.Reader) <-chan *Payload {
	payloadCh := make(chan *Payload)
	return payloadCh
}

func parse0(r io.Reader, ch chan<- *Payload) {
	reader := bufio.NewReader(r)
	for {
		header, err := reader.ReadBytes('\n')
		// reader 发生的错误告知 channel，并且停止解析
		if err != nil {
			goto ErrEnd
		}
		header = bytes.TrimSuffix(header, []byte{'\r', '\n'})
		switch header[0] {
		case '*':
			count, err := strconv.ParseInt(string(header[1:]), 10, 64)
			// 协议错误认为是客户端请求有问题，应该主动关闭连接，并且在返回中明确告知客户端
			if err != nil {
				protocolError(ch, "illegal array header: "+string(header[1:]))
				return
			}
			var command [][]byte
			for i := 0; i < int(count); i++ {
				item, err := reader.ReadBytes('\n')
				if err != nil {
					goto ErrEnd
				}
				item = bytes.TrimSuffix(item, []byte{'\r', '\n'})
				if item[0] != '$' {
					protocolError(ch, "the next line of * must start $, not "+string(item[0]))
					return
				}
				length, err := strconv.ParseInt(string(item[1:]), 10, 64)
				if err != nil {
					protocolError(ch, "parse to int failed: "+string(item[1:]))
					return
				}
				buf := make([]byte, length+int64(len(protocol.CLCR)))
				_, err = io.ReadFull(reader, buf)
				if err != nil {
					goto ErrEnd
				}
				buf = bytes.TrimSuffix(buf, []byte{'\r', '\n'})
				command = append(command, buf)
			}
			ch <- &Payload{
				Data: protocol.MultiBulk{
					Items: command,
				},
			}
		}
	ErrEnd:
		ch <- &Payload{Err: err}
		close(ch)
		return
	}
}

func protocolError(ch chan<- *Payload, msg string) {
	err := protocol.ProtocolErr{Msg: msg}
	ch <- &Payload{Err: err}
}
