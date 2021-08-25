package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry1(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry2(t *testing.T) {
	a := 7 // 0111
	t.Log(a&Readable == Readable) // true
	t.Log(a&Writable == Writable) // true
	t.Log(a&Executable == Executable) // true
}
