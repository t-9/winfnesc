/*
Package winfnesc is windows file name escaper.
*/
package winfnesc

import (
	"testing"
)

func TestEscape(t *testing.T) {
	expected := "￥　／　：　＊　？　”　＜　＞　｜"
	actueal := Escape(`\　/　:　*　?　"　<　>　|`)

	if expected != actueal {
		t.Fatalf("failed test\nexpected: %#v\nactual: %#v", expected, actueal)
	}
}
