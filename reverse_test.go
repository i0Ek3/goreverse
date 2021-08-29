package goreverse

import (
    "testing"

    "github.com/i0Ek3/asrt"
)

func TestReverseWithInterval(t *testing.T) {
    b := "hello"
    got := ReverseWithInterval([]byte(b), 0, 4)
    want := string(b)

    asrt.Equal(got, want)
}

func TestReverse(t *testing.T) {
    res := []int{1, 2, 3, 4, 5}
    got := Reverse(res)
    want := []int{5, 4, 3, 2, 1}

    asrt.Equal(got, want)
}

/*
func TestReverseAny(t *testing.T) {
    str := "hello"
    got := ReverseAny(str, 0, 4)
    want := "olleh"

    asrt.Equal(got, want)
}
*/
