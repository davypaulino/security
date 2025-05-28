package main

type CipherMessage struct {
	Message string
}

func encryptFactor(start rune, c rune, factor rune) string {
	return string(( (c - start + factor) % 26 + start))
}

func decryptFactor(start rune, c rune, factor rune) string {
	return string(( (c - start - factor + 26) % 26 + start))
}

func abs(number int) int {
	if (number >= 0) {
		return number
	}
	return number * -1
}