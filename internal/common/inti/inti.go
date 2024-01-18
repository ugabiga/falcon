package inti

import "strconv"

func CountZeros(n int) int {
	s := strconv.Itoa(n)
	count := 0
	for _, char := range s {
		if char == '0' {
			count++
		}
	}
	return count
}
