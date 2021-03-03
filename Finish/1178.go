package main

import "fmt"

func check(w string, p string) bool {
	pbyte := []byte(p)
	wbyte := []byte(w)
	pm := make(map[byte]int)
	wm := make(map[byte]int)

	for _, v := range wbyte {
		wm[v] = 1
	}
	if _, ok := wm[pbyte[0]]; !ok {
		return false
	}
	for _, v := range pbyte {
		pm[v] = 1
	}
	for _, v := range wbyte {
		if _, ok := pm[v]; !ok {
			return false
		}
	}
	return true
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	ret := make([]int, len(puzzles))
	for i, vx := range puzzles {
		for _, vy := range words {
			if check(vy, vx) {
				ret[i] += 1
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(findNumOfValidWords([]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}, []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}))
	fmt.Println(check("beefed", "abcdefg"))
}
