package boot

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	s1, err3 := Cache().Get("test-key1")
	fmt.Println(s1, err3)

	err1 := Cache().Set("test-key", []byte(""))
	s, err2 := Cache().Get("test-key")
	fmt.Println(err1, s, err2)
}
