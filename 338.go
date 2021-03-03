package main

import "fmt"

func itob_count(num int) (ret int) {
	ret = 0
	for num > 0 {
		num = num & (num - 1)
		ret++
	}
	return
}
func countBits(num int) []int {
	s := []int{}
	for i := 0; i <= num; i++ {
		s = append(s, itob_count(i))
	}
	return s
}

func countBits2(num int) []int {
	s := make([]int, num+1)
	for i := 0; i <= num; i++ {
		if i%2 == 0 {
			s[i] = s[i/2]
		} else {
			s[i] = s[i/2] + 1
		}
	}
	return s
}
func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
	fmt.Println(countBits2(2))
	fmt.Println(countBits2(5))
}
