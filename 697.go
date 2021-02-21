package main

import (
	"fmt"
	"sort"
)

type myStruct struct {
	count int
	left  int
	right int
}
type SortMyStruct []*myStruct

func (s SortMyStruct) Len() int { return len(s) }
func (s SortMyStruct) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortMyStruct) Less(i, j int) bool {
	return s[i].count > s[j].count
}

func findShortestSubArray(nums []int) int {
	var m = make(map[int]*myStruct)
	for index, v := range nums {
		//fmt.Println("v is ",v)
		if _, ok := m[v]; !ok {
			m[v] = &myStruct{
				count: 1,
				left:  index,
				right: index,
			}
		} else {
			m[v].count++
			m[v].right = index
		}
	}
	s := []*myStruct{}
	for _, v := range m {
		s = append(s, v)
		//fmt.Println(k,v.count,v.left,v.right)
	}
	sort.Sort(SortMyStruct(s))
	//fmt.Println("sorted")
	//for _,v:=range s{
	//	fmt.Println(v.count)
	//}
	mlen := s[0].count
	mret := 100000
	for _, v := range s {
		if v.count == mlen && v.right-v.left < mret {
			mret = v.right - v.left + 1
		}
	}
	//fmt.Println("mret is ",mret)
	return mret
}

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))

	a := []int{1, 1111, 4, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)

	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
