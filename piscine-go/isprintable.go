package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	for _, s := range str {
		if !(s >= 31) {
			return false
		}
	}
	return true
}
