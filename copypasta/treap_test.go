package copypasta

import (
	"fmt"
	"testing"
)

func Test_treap(t_ *testing.T) {
	// 使用 https://www.luogu.org/problem/P3369 来测试

	//_x = uint(1)
	t := newTreap()
	for i := 1; i < 100; i++ {
		t.put(tpKeyType(i), 1)
		fmt.Println(t)
	}
}
