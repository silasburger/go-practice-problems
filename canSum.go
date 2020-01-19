package main

import "fmt"

type intHolder struct {
	ints []int
}

func (i *intHolder) hasCompliment(target int) bool {
	intSet := make(map[int]bool)
	for _, val := range i.ints {
		intSet[target-val] = true
		if intSet[val] == true {
			return true
		}
	}
	return false
}

func main() {
	ints := intHolder{ints: []int{1, 2, 3, 4}}
	fmt.Println(ints.hasCompliment(7))
}
