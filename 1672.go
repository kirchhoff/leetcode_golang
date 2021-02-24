package main

import "fmt"

func maximumWealth(accounts [][]int) int {

	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			accounts[i][0] += accounts[i][j]
		}

		if accounts[i][0] > accounts[0][0] {
			accounts[0][0] = accounts[i][0]
		}
	}
	return accounts[0][0]

}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
}
