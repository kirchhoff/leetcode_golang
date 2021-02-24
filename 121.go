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

func maxProfit2(prices []int) (ret int) {
	ret = 0
	mxp := 0
	mi := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-mi > mxp {
			mxp = prices[i] - mi
		}
		if prices[i] < mi {
			mi = prices[i]
		}
	}
	ret = mxp
	return
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 2, 1}))

	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit2([]int{7, 6, 4, 2, 1}))
}
