package main

import "github.com/01-edu/z01"

func main() {
	QuadA(5, 3)
}

func QuadA(x, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 && j == 0 {
				z01.PrintRune('o')
			} else if i == 0 && j == x-1 {
				z01.PrintRune('o')
			} else if i == y-1 && j == 0 {
				z01.PrintRune('o')
			} else if i == y-1 && j == x-1 {
				z01.PrintRune('o')
			} else if j == 0 || j == y-1 {
				z01.PrintRune('-')
			} else if j == 0 || j == x-1 {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
