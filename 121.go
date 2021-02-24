package main

import (
	"fmt"
)

func maxProfit(prices []int) (ret int) {
	ret = 0
	for i := 0; i < len(prices); {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > ret {
				ret = prices[j] - prices[i]
			}
		}
		k := i + 1
		for ; k < len(prices); k++ {
			if prices[k] < prices[i] {
				break
			}
			//fmt.Println("xxxx")
		}
		i = k
	}
	return
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 2, 1}))
}
