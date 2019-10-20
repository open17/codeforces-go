package copypasta

// 耗时大约是红黑树的 1.3 倍

import (
	. "fmt"
	"time"
)

var _x = uint(time.Now().UnixNano())

// https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
// https://en.wikipedia.org/wiki/Xorshift
func fastRand() uint {
	_x ^= _x << 13
	_x ^= _x >> 17
	_x ^= _x << 5
	return _x
}

type tpKeyType int   // *custom* 图方便可以全局替换
type tpValueType int // *custom* 图方便可以全局替换

type tpNode struct {
	lr       [2]*tpNode
	priority uint // max heap
	sz       int
	msz      int
	key      tpKeyType
	value    tpValueType
}

// d=0: left
// d=1: right
func (o *tpNode) rotate(d int) *tpNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	// x.msz = o.msz; o.pushUp()
	o.pushUp()
	x.pushUp()
	return x
}

func (o *tpNode) pushUp() {
	sz := 1
	msz := int(o.value)
	if ol := o.lr[0]; ol != nil {
		sz += ol.sz
		msz += ol.msz
	}
	if or := o.lr[1]; or != nil {
		sz += or.sz
		msz += or.msz
	}
	o.sz = sz
	o.msz = msz
}

type treap struct {
	root       *tpNode
	comparator func(a, b tpKeyType) int
}

func newTreap() *treap {
	// 设置如下返回值是为了方便使用 tpNode 中的 lr 数组
	return &treap{comparator: func(a, b tpKeyType) int {
		if a < b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	}}
}

func (t *treap) _put(o *tpNode, key tpKeyType, value tpValueType) *tpNode {
	if o == nil {
		return &tpNode{priority: fastRand(), sz: 1, msz: 1, key: key, value: value}
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key, value)
		if o.lr[cmp].priority > o.priority {
			o = o.rotate(cmp ^ 1)
		}
	} else {
		o.value = value
		//o.value += value
	}
	o.pushUp()
	return o
}

func (t *treap) put(key tpKeyType, value tpValueType) { t.root = t._put(t.root, key, value) }

func (t *treap) _delete(o *tpNode, key tpKeyType) *tpNode {
	if o == nil {
		return nil
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._delete(o.lr[cmp], key)
	} else {
		//if o.value > 1 {
		//	o.value--
		//} else {
		if o.lr[1] == nil {
			return o.lr[0]
		}
		if o.lr[0] == nil {
			return o.lr[1]
		}
		// o 有两颗子树，先把优先级高的子树旋转到根，然后递归在另一颗子树中删除 o
		cmp2 := 0
		if o.lr[0].priority > o.lr[1].priority {
			cmp2 = 1
		}
		o = o.rotate(cmp2)
		o.lr[cmp2] = t._delete(o.lr[cmp2], key)
		//}
	}
	o.pushUp()
	return o
}

func (t *treap) delete(key tpKeyType) { t.root = t._delete(t.root, key) }

//

func (o *tpNode) String() string {
	var s string
	if o.value == 1 {
		s = Sprintf("%v", o.key)
	} else {
		s = Sprintf("%v(%v)", o.key, o.value)
	}
	s += Sprintf("[sz:%d,msz:%d]", o.sz, o.msz)
	return s
}

func (o *tpNode) draw(prefix string, isTail bool, str *string) {
	if o.lr[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		o.lr[1].draw(newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += o.String() + "\n"
	if o.lr[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		o.lr[0].draw(newPrefix, true, str)
	}
}

func (t *treap) String() string {
	if t.root == nil {
		return "BST (empty)\n"
	}
	str := "BST\n"
	t.root.draw("", true, &str)
	return str
}
