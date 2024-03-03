package main

import "fmt"

func main() {
	as := [][]int{{2, 4, 6, 2, 5}, {5, 1, 1, 5}, {1, 2, 3, 4, 5, 6}}
	for _, a := range as {
		fmt.Println("Maximum non-adjacent sum of numbers:", a)
		fmt.Println("is =>", maxNonadjecentSum(a))
	}
}

func maxNonadjecentSum(a []int) int {
	if len(a) < 1 {
		return 0
	}
	maxSumWithout, maxSumWith := 0, a[0]
	for _, x := range a[1:] {
		maxSumWith, maxSumWithout = maxSumWithout+x, max(maxSumWith, maxSumWithout)
	}
	return max(maxSumWith, maxSumWithout)
}
