package parser

import (
	"strings"
	"testing"

	"github.com/liamgheng/godis/redis/protocol"
	"github.com/stretchr/testify/assert"
)

func TestParser0(t *testing.T) {
	payloadCh := make(chan *Payload)
	r := strings.NewReader("*2\r\n$4\r\nLLEN\r\n$6\r\nmylist\r\n")
	go parse0(r, payloadCh)
	payload := <-payloadCh

	expect := &Payload{
		Data: protocol.MultiBulk{
			Items: [][]byte{[]byte("LLEN"), []byte("mylist")},
		},
	}

	assert.EqualValues(t, expect, payload)
}
