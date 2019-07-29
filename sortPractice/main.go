package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}
func (p people) Less(i, j int) bool {
	if p[i] > p[j] {
		return true
	} else {
		return false
	}
}
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Stable(studyGroup)
	fmt.Println(studyGroup)
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	fmt.Println(s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Slice(n, func(i, j int) bool { return n[i] > n[j] })
	fmt.Println(n)

}
