package main

import (
	"slices"
)

func anyTwoSumToK(a []int, k int) bool {
	has := map[int]bool{}
	for _, x := range a {
		if has[k-x] {
			return true
		}
		has[x] = true
	}
	return false
}
func anyTwoSumToKIndHashmap(a []int, k int) []int {
	indices := map[int]int{}
	for i, x := range a {
		ind, has := indices[k-x]
		if has {
			return []int{ind, i}
		}
		indices[x] = i
	}

	return []int{0, 0}
}

func anyTwoSumToKIndSort(a []int, k int) bool {
	slices.Sort(a)
	sum := 0
	for i, j := 0, len(a)-1; i < j; {
		sum = a[i] + a[j]
		if sum == k {
			return true
		} else if sum < k {
			i++
		} else {
			j--
		}
	}
	return false
}

func main() {

}
