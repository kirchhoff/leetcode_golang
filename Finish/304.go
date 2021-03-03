package main

import "fmt"

type NumMatrix struct {
	nums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nums := make([][]int, len(matrix))
	for i, row := range matrix {
		nums[i] = make([]int, len(row)+1)
		for j, v := range row {
			nums[i][j] = v
		}
	}
	return NumMatrix{nums: nums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (ret int) {
	ret = 0
	for i := row1; i <= row2; i++ { //2,4
		for j := col1; j <= col2; j++ { //1 8
			ret += this.nums[i][j]
		}
	}
	return
}
func main() {

	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	obj := Constructor(matrix)
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
