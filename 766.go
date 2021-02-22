package main

import "fmt"

func getAndCheck(x, y int, matrix [][]int, row, column int) bool {
	nums := []int{}

	for x < row && y < column {
		nums = append(nums, matrix[x][y])
		x++
		y++
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if nums[0]*len(nums) != sum {
		return false
	}
	return true
}

func isToeplitzMatrix(matrix [][]int) bool {
	var x = len(matrix)
	var y = len(matrix[0])
	//test line one
	for j := 0; j < y; j++ {
		if getAndCheck(0, j, matrix, x, y) == false {
			return false
		}
	}
	for i := 0; i < x; i++ {
		if getAndCheck(i, 0, matrix, x, y) == false {
			return false
		}
	}
	return true
}

func isToeplitzMatrix2(matrix [][]int) bool {
	var x = len(matrix)
	var y = len(matrix[0])
	if x == 1 && y == 1 {
		return true
	}
	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}

	return true
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(matrix))
	fmt.Println(isToeplitzMatrix2(matrix))
	matrix = [][]int{{1, 2}, {2, 2}}
	fmt.Println(isToeplitzMatrix(matrix))
	fmt.Println(isToeplitzMatrix2(matrix))
}

/*输入：matrix = [[1,2,3,4],[5,1,2,3],[9,5,1,2]]
输出：true
解释：
在上述矩阵中, 其对角线为:
"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]"。
各条对角线上的所有元素均相同, 因此答案是 True 。
。*/
