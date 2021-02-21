package main

import "fmt"

func minSubArrayLen(target int, nums []int) (ret int) {
	ret = 100001
	var l, r = 0, 0
	sum := 0
	for r < len(nums) {
		for r < len(nums) && sum < target {
			sum += nums[r]
			r++
		}
		for l < len(nums) && sum >= target {
			sum -= nums[l]
			l++
		}
		ret = min(ret, r-l+1)
		//r++
	}
	if ret == len(nums)+1 {
		ret = 0
	}
	return
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
