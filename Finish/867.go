package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, m)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			s[j][i] = matrix[i][j]
		}
	}
	return s
}

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
