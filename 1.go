package main

import "fmt"

func twoSum(nums []int,target int)[]int{
	lenOfNums := len(nums)
	for i,vx:=range nums{
		for j:=i+1;j<lenOfNums;j++{
			if nums[j]+vx==target {
				return []int{i, j}
			}
		}
	}
	return nil
}


func main()  {
	nums:=[...]int{3,2,4}
	fmt.Println(twoSum(nums[:],6))
}
