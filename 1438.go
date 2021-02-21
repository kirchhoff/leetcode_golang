package main

import "fmt"

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
