package main

import "fmt"

func lengthOfLongestSubstring(s string) (ret int) {
	ret = 0
	sbyte := []byte(s)
	var l, r int
	//fmt.Println(l)
	m := make(map[byte]int)
	//ts:=[]byte{}
	for r < len(s) {
		//fmt.Println("current is ",sbyte[r])
		if _, ok := m[sbyte[r]]; !ok {
			m[sbyte[r]] = 1
			//fmt.Println("set!")
			//ts=append(ts,sbyte[r])
		} else {
			for sbyte[l] != sbyte[r] {
				delete(m, sbyte[l])
				l++
			}
			if sbyte[l] == sbyte[r] {
				m[sbyte[l]] = 1

			}
			l++
		}
		r++
		ret = maxf(ret, r-l)
	}
	return
}
func maxf(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
