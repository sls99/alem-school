package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 9223372036854775807 {
		return 0
	}
	fact := 1
	for i := 1; i <= nb; i++ {
		fact *= i
	}
	return fact
}
