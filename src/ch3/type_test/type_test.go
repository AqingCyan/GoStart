package type_test

import "testing"

type MyInt int64 // 类型别名


func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64

	// b = a 无法进行赋值，需要强制类型转换
	b = int64(a)
	t.Log(a, b)

	// c = b 无法进行赋值，即使 MyInt 本身就是int64
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}
