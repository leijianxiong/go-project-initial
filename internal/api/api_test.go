package api

import (
	"log"
	"testing"
)

func TestByte(t *testing.T) {
	b := make([]byte, 2)
	b[0] = 'a'
	b[1] = 'b'
	b = b[:3]
	log.Println("b: ", string(b))
}
