package main

import "fmt"

func firstMissingPositiveInt(a []int) int {
	n := len(a)

	for i := 0; i < n; i++ {
		for a[i] > 0 && a[i] <= n && a[i] != i+1 {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}

	for i := 0; i < n; i++ {
		if a[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	fmt.Println(firstMissingPositiveInt([]int{1, 2, 6, 3, 5, 4}))
}
