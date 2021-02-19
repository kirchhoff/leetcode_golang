package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r1 := reverseList(l1)
	r2 := reverseList(l2)

	for l1 != nil || l2 != nil {

	}
	for i, p := 0, l1; p != nil; p = p.Next {

		v1 = v1 + int(math.Pow10(i))*p.Val
		//fmt.Println("v1",v1)
		i = i + 1
	}
	for i, p := 0, l2; p != nil; p = p.Next {

		v2 = v2 + int(math.Pow10(i))*p.Val
		i = i + 1
	}
	fmt.Println(v1, v2)
	v3 := v1 + v2
	/*fmt.Println(v3)
	s3 := []rune(strconv.Itoa(v3))
	for j := 0; j < len(s3)/2; j++ {
		s3[j], s3[len(s3)-j-1] = s3[len(s3)-j-1], s3[j]
	}
	fmt.Println("s3 is ",string(s3))
	ret,_:=strconv.ParseInt(string(s3),10,32)*/

	return iToListNode(v3)

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
	//n1 := 1000000000000000001
	//l1 := iToListNode(n1)
	n2 := 564
	l2 := iToListNode(n2)
	printList(l2)
	fmt.Println()
	printList(reverseList(l2))
	//addTwoNumbers(l1,l2)

	//342 + 465 = 807.

}
