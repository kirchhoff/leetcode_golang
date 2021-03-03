package main

import "fmt"

//
//import "math/rand"
//
//type node struct {
//	ch       [2]*node
//	priority int
//	key      int
//	cnt      int
//}
//
//func (o *node) cmp(b int) int {
//	switch {
//	case b < o.key:
//		return 0 //小 左边
//	case b > o.key:
//		return 1
//	default:
//		return -1
//	}
//}
//
//func (o *node) rotate(d int) *node {   //d = 0左转   d=1 右转
//	x := o.ch[d^1]
//	o.ch[d^1] = x.ch[d]
//	x.ch[d] = o
//	return x
//}
//
//type treap struct {
//	root *node
//}
//
//func (t *treap) ins(o *node, key int) *node {
//	if o == nil {
//		return &node{priority: rand.Int(), key: key, cnt: 1}
//	}
//	if d := o.cmp(key); d >= 0 {
//		o.ch[d] = t.ins(o.ch[d], key)
//		if o.ch[d].priority > o.priority {
//			o = o.rotate(d ^ 1)
//		}
//	} else { //相等的插入
//		o.cnt++
//	}
//	return o
//}
//
//func (t *treap) del(o *node, key int) *node {
//	if o == nil {
//		return nil
//	}
//	if d := o.cmp(key); d >= 0 {
//		o.ch[d] = t.del(o.ch[d], key)
//	} else {
//		if o.cnt > 1 {
//			o.cnt--
//		} else {
//			if o.ch[1] == nil {
//				return o.ch[0]
//			}
//			if o.ch[0] == nil {
//				return o.ch[1]
//			}
//			d = 0
//			if o.ch[0].priority > o.ch[1].priority {
//				d = 1
//			}
//			o = o.rotate(d)
//			o.ch[d] = t.del(o.ch[d], key)
//		}
//	}
//	return o
//}
//
//func (t *treap) insert(key int) {
//	t.root = t.ins(t.root, key)
//}
//
//func (t *treap) delete(key int) {
//	t.root = t.del(t.root, key)
//}
//
//func (t *treap) min() (min *node) {
//	for o := t.root; o != nil; o = o.ch[0] {
//		min = o
//	}
//	return
//}
//
//func (t *treap) max() (max *node) {
//	for o := t.root; o != nil; o = o.ch[1] {
//		max = o
//	}
//	return
//}

func longestSubarray2(nums []int, limit int) (ans int) {
	t := &treap{}
	left := 0
	for right, v := range nums {
		t.insert(v)
		for t.max().key-t.min().key > limit {
			t.delete(nums[left])
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func abs(n int) (ret int) {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func longestSubarray(nums []int, limit int) (ret int) {
	var maxQ []int
	var minQ []int
	ret = 0
	var l, r = 0, 0
	for r < len(nums) {
		//fmt.Println("nums[r] is ",nums[r])
		for len(maxQ) > 0 && nums[r] > maxQ[len(maxQ)-1] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		for len(minQ) > 0 && nums[r] < minQ[len(minQ)-1] {
			minQ = minQ[:len(minQ)-1]
		}
		maxQ = append(maxQ, nums[r])
		//fmt.Println(maxQ)
		minQ = append(minQ, nums[r])
		//fmt.Println(minQ)
		for len(maxQ) > 0 && len(minQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[l] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			if nums[l] == minQ[0] {
				minQ = minQ[1:]
			}
			l++
		}
		ret = max(ret, r-l+1)
		r++
	}

	return
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarray([]int{1, 5, 6, 7, 8, 10, 6, 5, 6}, 4))

}
