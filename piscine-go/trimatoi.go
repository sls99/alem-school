package piscine

func TrimAtoi(s string) int {
	Run := []rune(s)
	counter := 0
	sing := 1

	for i := 0; i < len(s); i++ {
		if counter == 0 && Run[i] == 45 {
			sing = -1
		}
		if Run[i] > 47 && Run[i] < 58 {
			counter = counter*10 + int(Run[i]-48)
		}
	}

	return counter * sing
}
