package main

import "fmt"

func main() {
	const msg = "111"
	fmt.Println(numOfDecodings(msg))
}

func numOfDecodings(msg string) int {
	n := len(msg)

	msgr := []rune(msg)

	memo := make([]int, n+1)
	if validOne(msgr[0]) {
		memo[1]++
	}
	if n == 1 {
		return memo[1]
	}
	if validOne(msgr[1]) {
		if memo[1] > 0 {
			memo[2]++
		}
	}
	if validTwo(msgr[0:2]) {
		memo[2]++
	}

	for m := 3; m <= n; m++ {
		i := m - 1
		if validOne(msgr[i]) {
			if memo[m-1] > 0 {
				memo[m] += memo[m-1]
			}
		}
		if validTwo(msgr[i-1 : i+1]) {
			if memo[m-2] > 0 {
				memo[m] += memo[m-2]
			}
		}
	}

	return memo[n]
}

func validTwo(binome []rune) bool {
	a := binome[0]
	b := binome[1]
	if a == '1' && 0 <= b && b <= '9' {
		return true
	}
	if a == '2' && 0 <= b && b <= '6' {
		return true
	}
	return false
}

func validOne(c rune) bool {
	return c >= '1' && c <= '9'
}
