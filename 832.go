package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		j := 0
		for j = 0; j < len(A[i])/2; j++ {
			A[i][j], A[i][len(A[i])-j-1] = A[i][len(A[i])-j-1], A[i][j]
			A[i][j] ^= 1
			A[i][len(A[i])-j-1] ^= 1
		}
		if j == len(A[i])/2 && len(A[i])%2 == 1 {
			A[i][len(A[i])-j-1] ^= 1
		}
	}
	return A
}

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
	//[[1,0,0],[0,1,0],[1,1,1]]
	//[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
}
