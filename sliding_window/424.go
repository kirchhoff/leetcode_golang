package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	ch       [2]*node
	priority int
	key      int
	cnt      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.key:
		return 0 //小 左边
	case b > o.key:
		return 1 //大 右边
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node { //d = 0左转   d=1 右转
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

type treap struct {
	root *node
}

func (t *treap) ins(o *node, key int) *node {
	if o == nil {
		return &node{priority: rand.Int(), key: key, cnt: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.ch[d] = t.ins(o.ch[d], key)
		if o.ch[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else { //相等的插入
		o.cnt++
	}
	return o
}

func (t *treap) del(o *node, key int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.ch[d] = t.del(o.ch[d], key)
	} else {
		if o.cnt > 1 {
			o.cnt--
		} else {
			if o.ch[1] == nil {
				return o.ch[0]
			}
			if o.ch[0] == nil {
				return o.ch[1]
			}
			d = 0
			if o.ch[0].priority > o.ch[1].priority {
				d = 1
			}
			o = o.rotate(d)
			o.ch[d] = t.del(o.ch[d], key)
		}
	}
	return o
}

func (t *treap) insert(key int) {
	t.root = t.ins(t.root, key)
}

func (t *treap) delete(key int) {
	t.root = t.del(t.root, key)
}

func (t *treap) min() (min *node) {
	for o := t.root; o != nil; o = o.ch[0] {
		min = o
	}
	return
}

func (t *treap) max() (max *node) {
	for o := t.root; o != nil; o = o.ch[1] {
		max = o
	}
	return
}
func characterReplacement(s string, k int) (ret int) {

	if len(s) == 0 {
		return 0
	}
	var l, r = 0, 1
	ret = 0
	m := make(map[byte]int)
	sbyte := []byte(s)
	//fmt.Println("sbyte is ",sbyte)
	t := &treap{}
	t.insert(1)
	m[sbyte[l]] = 1
	for r < len(s) {
		if v, ok := m[sbyte[r]]; !ok {
			t.insert(1)
			m[sbyte[r]] = 1
		} else {
			t.delete(v)
			m[sbyte[r]] = v + 1
			t.insert(m[sbyte[r]])
			//fmt.Println("current map is ",m)
		}
		for t.max().key+k < (r - l + 1) {
			t.delete(m[sbyte[l]])
			m[sbyte[l]] -= 1
			t.insert(m[sbyte[l]])
			l++
		}
		if t.max().key+k >= (r - l + 1) {
			ret = maxij(ret, r-l+1)
		}
		r++
	}
	return
}
func maxij(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {

	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
	fmt.Println(characterReplacement("AAAA", 0))
	fmt.Println(characterReplacement("ABCDE", 1))
}

//s = "AABABBA", k = 1   ->4
