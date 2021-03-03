package main

import (
	"fmt"
)

func maxSatisfied(customers []int, grumpy []int, X int) (ret int) {
	var l, r = 0, 0
	ret = 0
	ts := 0
	if len(customers) == 1 {
		ret = customers[0]
		return
	}
	for i := 0; i < X; i++ {
		ts += customers[i]
	}
	ret = ts
	left := 0
	r = X
	right := 0
	for j := r; j < len(customers); j++ {
		if grumpy[j] == 0 {
			right += customers[j]
		}

	}
	if ts+left+right > ret {
		ret = ts + left + right
	}
	//fmt.Println("l,r,ts",l,r,ts)
	for r < len(customers) {
		//fmt.Println("r is ",r)
		ts += customers[r]
		ts -= customers[l]
		if grumpy[l] == 0 {
			left += customers[l]
		}

		l++
		if grumpy[r] == 0 {
			right -= customers[r]
		}
		if ts+left+right > ret {
			ret = ts + left + right
		}
		r++
	}
	return
}

//
//customers = [1,0,1,2,1,1,7,5],
//  grumpy = [0,1,0,1,0,1,0,1], X = 3
//16
func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3) == 16)
	fmt.Println(maxSatisfied([]int{5, 8}, []int{0, 1}, 1) == 13)
	fmt.Println(maxSatisfied([]int{3}, []int{1}, 1) == 3)
	fmt.Println(maxSatisfied([]int{3, 7}, []int{0, 0}, 2) == 10)
	fmt.Println(maxSatisfied([]int{10, 4}, []int{0, 1}, 2) == 14)
	fmt.Println(maxSatisfied([]int{4, 10, 10}, []int{1, 1, 0}, 2) == 24) //24
	fmt.Println(maxSatisfied([]int{10, 1, 7}, []int{0, 0, 0}, 2) == 18)
}
