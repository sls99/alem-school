package piscine

func ReverseString(s string) string {
	runes := []rune(s)
	size := len(runes)
	for i := 0; i < size/2; i++ {
		runes[size-i-1], runes[i] = runes[i], runes[size-i-1]
	}
	return string(runes)
}
