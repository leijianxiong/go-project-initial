package toolutil

import (
	"fmt"
	"testing"
)

func TestMustNoErrorf(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			s := fmt.Sprintf("%s", p)
			if s != "fad 1, a, err1" {
				t.Error()
			}
		}
	}()
	MustNoErrorf("fad %v, %v, %s", 1, "a", fmt.Errorf("err1"))
}

func TestMustNoErrorfForNoErrVar(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			s := fmt.Sprintf("%s", p)
			if s != "id missing" {
				t.Error()
			}
		}
	}()
	MustNoErrorf("id missing")
}