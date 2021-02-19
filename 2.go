package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}
func reverseList(l *ListNode) *ListNode {
	tail := l
	if tail != nil {
		for tail.Next != nil {
			tail = tail.Next
		}
		//fmt.Println(tail.Val)
		for l != tail {
			t := l.Next
			l.Next = tail.Next
			tail.Next = l
			l = t
		}
		return tail
	} else {
		return nil
	}
}

func addTwoNumbers(r1 *ListNode, r2 *ListNode) *ListNode {

	flag := 0
	var ret = new(ListNode)
	var iter *ListNode
	//var once=true
	for r1 != nil || r2 != nil {
		iter = new(ListNode)
		if r1 != nil && r2 != nil {
			iter.Val = (r1.Val + r2.Val + flag) % 10
			//fmt.Println("current iter",iter.Val)
			flag = (r1.Val + r2.Val + flag) / 10
			//fmt.Println("flag is ",flag)
			r1 = r1.Next
			r2 = r2.Next
		} else if r1 == nil {
			iter.Val = (r2.Val + flag) % 10
			flag = (r2.Val + flag) / 10
			r2 = r2.Next
			//fmt.Println("current iter",iter.Val)
		} else {
			iter.Val = (r1.Val + flag) % 10
			flag = (r1.Val + flag) / 10
			r1 = r1.Next
			//fmt.Println("current iter",iter.Val)
		}

		iter.Next = ret.Next
		//fmt.Println("ret next isss ",ret.Next)
		ret.Next = iter

		//printList(ret.Next)
	}
	if flag != 0 {
		iter = new(ListNode)
		iter.Val = flag
		iter.Next = ret.Next
		//fmt.Println("ret next isss ",ret.Next)
		ret.Next = iter

	}
	//printList(ret.Next)
	//printList(ret.Next)
	//printList(reverseList(ret.Next))

	return reverseList(ret.Next)

}

func addTwoNumbers2(r1 *ListNode, r2 *ListNode) *ListNode {

	flag := 0
	var ret = new(ListNode)
	fin := ret
	var iter *ListNode
	for r1 != nil || r2 != nil {
		iter = new(ListNode)
		if r1 != nil && r2 != nil {
			iter.Val = (r1.Val + r2.Val + flag) % 10
			//fmt.Println("current iter",iter.Val)
			flag = (r1.Val + r2.Val + flag) / 10
			//fmt.Println("flag is ",flag)
			r1 = r1.Next
			r2 = r2.Next
		} else if r1 == nil {
			iter.Val = (r2.Val + flag) % 10
			flag = (r2.Val + flag) / 10
			r2 = r2.Next
			//fmt.Println("current iter",iter.Val)
		} else {
			iter.Val = (r1.Val + flag) % 10
			flag = (r1.Val + flag) / 10
			r1 = r1.Next
			//fmt.Println("current iter",iter.Val)
		}

		ret.Next = iter
		//fmt.Println("ret next isss ",ret.Next)
		ret = iter

		//printList(ret.Next)
	}
	if flag != 0 {
		iter = new(ListNode)
		iter.Val = flag
		ret.Next = iter
		ret = iter

	}
	//printList(ret.Next)
	//printList(ret.Next)
	//printList(reverseList(ret.Next))

	return fin.Next

}

func iToListNode(v int) *ListNode {
	var ret *ListNode
	var l *ListNode
	if v == 0 {
		ret = new(ListNode)
		return ret
	}
	for v != 0 {
		if l == nil {
			l = new(ListNode)
			ret = l
		}
		l.Val = v % 10
		if v/10 != 0 {
			l.Next = new(ListNode)
		}
		v = v / 10
		l = l.Next
	}

	return ret
}
func main() {
	n1 := 73
	l1 := iToListNode(n1)
	printList(l1)
	n2 := 29
	l2 := iToListNode(n2)
	printList(l2)
	fmt.Println()
	//printList(reverseList(l2))
	//printList(addTwoNumbers2(l1, l2))
	printList(addTwoNumbers3(l1, l2))

	//342 + 465 = 807.

}

func addTwoNumbers3(l1, l2 *ListNode) (ret *ListNode) {
	carry := 0
	var tail *ListNode
	var t *ListNode
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		summ := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		if ret == nil {
			ret = &ListNode{
				Val: summ,
			}
			tail = ret
		} else {
			t = &ListNode{Val: summ}
			tail.Next = t
			tail = tail.Next
		}
	}
	if carry != 0 {
		t = &ListNode{Val: carry}
		tail.Next = t
		tail = tail.Next
	}
	return ret
}
