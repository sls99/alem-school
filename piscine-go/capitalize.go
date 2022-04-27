package piscine

func Capitalize(s string) string {
	y := 0
	new := []rune(s)
	for range s {
		y++
	}
	for i := 0; i < y; i++ {
		if new[i] >= 'A' && new[i] <= 'Z' {
			new[i] = (new[i] + 32)
		}
		if new[0] >= 'a' && new[0] <= 'z' {
			new[0] = (new[0] - 32)
		}
		if i > 0 {
			if (new[i-1] > 'Z' && new[i-1] < 'a') || (new[i-1] < 'A' && new[i-1] > '9') || new[i-1] < '0' || new[i-1] > 'z' {
				if new[i] >= 'a' && new[i] <= 'z' {
					new[i] = (new[i] - 32)
				}
			}
		}
	}
	str := string(new)
	return str
}
