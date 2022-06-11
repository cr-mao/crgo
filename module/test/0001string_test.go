package test

import (
	"testing"
)

func TestString(t *testing.T) {
	s := "中"
	c := []rune(s)
	t.Logf("中的unicode是 %x ", c[0]) // 0x4e2d
	t.Logf("中的utf-8是 %x ", s)      //[0xe4 , 0xb8 ,0xad]
}
