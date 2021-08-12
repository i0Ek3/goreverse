package goreverse

import (
    "testing"

    "github.com/i0Ek3/asrt"
)

func TestReverse(t *testing.T) {
    b := "hello"
    got := Reverse([]byte(b), 0, 4)
    want := string(b)

    asrt.Equal(t, got, want)
}

/*
func TestReverseAny(t *testing.T) {
    str := "hello"
    got := ReverseAny(str, 0, 4)
    want := "olleh"

    asrt.Equal(t, got, want)
}
*/
