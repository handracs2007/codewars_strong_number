package main

import "fmt"

func factorial(n int) int {
	if n < 2 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func Strong(n int) string {
	var temp = n
	var sum = 0

	for temp > 0 {
		sum += factorial(temp % 10)
		temp /= 10
	}

	if n == sum {
		return "STRONG!!!!"
	} else {
		return "Not Strong !!"
	}
}

func main() {
	fmt.Println(Strong(1))
	fmt.Println(Strong(2))
	fmt.Println(Strong(145))
	fmt.Println(Strong(7))
	fmt.Println(Strong(93))
	fmt.Println(Strong(185))
}
