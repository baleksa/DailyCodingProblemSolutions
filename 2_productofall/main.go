package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 2, 1}
	fmt.Println(productOfAllNSpace(b))
	fmt.Println(productOfAllNSpace(a))
}

func productOfAll(a []int) []int {
	n := len(a)
	if n == 1 {
		return a
	}

	result := make([]int, n)

	result[0] = a[0]
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * a[i]
	}

	rprod := 1
	for i := n - 1; i > 0; i-- {
		result[i] = result[i-1] * rprod
		rprod *= a[i]
	}
	result[0] = rprod

	return result
}
func productOfAllNSpace(a []int) []int {
	n := len(a)
	if n == 1 {
		return a
	}

	result := make([]int, n)
	fstart := make([]int, n)
	fend := make([]int, n)

	fstart[0] = a[0]
	for i := 1; i < n; i++ {
		fstart[i] = fstart[i-1] * a[i]
	}

	fend[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		fend[i] = fend[i+1] * a[i]
	}

	result[0] = fend[1]
	result[n-1] = fstart[n-2]
	for i := 1; i < n-1; i++ {
		l := fstart[i-1]
		r := fend[i+1]
		result[i] = l * r
	}

	return result
}
