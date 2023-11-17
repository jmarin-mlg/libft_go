package main

import "unicode"

func ftAtoiImproved(str string) int {
	sign := 1
	strLen := len(str)
	result := 0

	// Skip leading whitespaces
	i := 0
	for i < strLen && unicode.IsSpace(rune(str[i])) {
		i++
	}

	// Handle sign
	if i < strLen && (str[i] == '+' || str[i] == '-') {
		if str[i] == '-' {
			sign = -1
		}

		i++
	}

	// Convert the number
	for i < strLen && unicode.IsDigit(rune(str[i])) {
		result = result*10 + int(str[i]-'0')

		i++
	}

	return result * sign
}
