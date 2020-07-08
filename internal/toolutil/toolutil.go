package toolutil

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
)

func PhoneMask(phone string, mask string) string {
	phoneBytes := []byte(phone)
	if mask == "" {
		mask = strings.Repeat("*", 4)
	}
	copy(phoneBytes[3:7], []byte(mask))
	return string(phoneBytes)
}

func MustNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func MustNoErrorf(format string, a ...interface{}) {
	var err error

	hasErrVar := false
	for _, v := range a {
		if _, ok := v.(error); ok {
			err = v.(error)
			hasErrVar = true
			break
		}
	}

	if !hasErrVar || err != nil {
		panic(fmt.Errorf(format, a...))
	}
}

func Sum(content string) (s string) {
	if content == "" {
		content = time.Now().String()
	}
	s = fmt.Sprintf("%x", md5.Sum([]byte(content)))
	return
}