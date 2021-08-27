package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	e := [...]int{1, 3, 2, 5}

	t.Log(a == b) // false 因为元素有不一致的
	t.Log(a == d) // true
	t.Log(a == e) // false 因为元素顺序不一致
	//t.Log(a == c) 无效运算: a == c(类型 [4]int 和 [5]int 不匹配)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	// 按位清零，清除它们的可读可写可执行
	a = a &^ Readable
	a = a &^ Executable
	a = a &^ Writable
	t.Log(a&Readable == Readable, a&Executable == Executable, a&Writable == Writable)
}
