package main

import "fmt"

// isPalindrom return true if n is a plainfrom number
func isPalindrom(n int) bool {
	s := fmt.Sprintf("%d", n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	max := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			n := i * j
			if isPalindrom(n) && n > max {
				max = n
			}
		}
	}
	fmt.Println(max)
}