package parser

import (
	"errors"
	"io"

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

// func parse0(r io.Reader, ch chan<- *Payload) {
// 	reader := bufio.NewReader(r)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err != nil {
// 			ch <- &Payload{Err: err}
// 			close(ch)
// 			return
// 		}
// 		line = strings.TrimSuffix(line, "\r\n")
// 		// switch line[0] {
// 		// case '+':
// 		// 	ch <- &Payload{
// 		// 		Data: protocol.MakeStatusReply(line[1:]),
// 		// 	}
// 		// case '-':
// 		// 	ch <- &Payload{
// 		// 		Data: protocol.MakeErrorReply(line[1:]),
// 		// 	}
// 		// case ':':
// 		// 	num, err := strconv.ParseInt(line[1:], 10, 64)
// 		// 	if err != nil {
// 		// 		ch <- &Payload{Err: err}
// 		// 		protocolError(line[1:], ch)
// 		// 		continue
// 		// 	}
// 		// 	ch <- &Payload{
// 		// 		Data: protocol.MakeIntReply(num),
// 		// 	}
// 		// case '$':
// 		// 	contentLen, err := strconv.ParseInt(line[1:], 10, 64)
// 		// 	if err != nil || contentLen < -1 {
// 		// 		protocolError("illegal bulk string header: " + line[1:], ch)
// 		// 		continue
// 		// 	} else if contentLen == -1 {

// 		// 	}
// 		// 	contentLine, err := reader.ReadString('\n')
// 		// 	if err != nil {
// 		// 		ch <- &Payload{Err: err}
// 		// 		close(ch)
// 		// 		return
// 		// 	}
// 		// 	contentLine = strings.TrimSuffix(contentLine, "\r\n")
// 		// 	if len(contentLine) != int(contentLen) {
// 		// 		// TODO
// 		// 		protocolError("tmp", ch)
// 		// 	}
// 		// }
// 	}
// }

func protocolError(msg string, ch chan<- *Payload) {
	err := errors.New("prototol error: " + msg)
	ch <- &Payload{Err: err}
}
