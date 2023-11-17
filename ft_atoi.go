package main

/*
** The function converts the initial portion of the string pointed to by str to
** int representation.
** ASCII 9 - 13 and 32 => Spaces
** ASCII 43 and 45     => Signs + and -
** ASCII 48 - 57       => Numbers 0 - 9
 */
func ftAtoi(str string) int {
	sign := 1
	signCount := 0
	strLen := ftStrlen(str)
	result := 0

	// Find the start of the number
	i := 0
	for i < strLen && (str[i] == 9 || (str[i] >= 10 && str[i] <= 13) || str[i] == 32) {
		i++
	}

	// Handle signs
	for i < strLen && (str[i] == 43 || str[i] == 45) {
		if str[i] == 45 {
			sign *= -1
		}

		i++

		if signCount++; signCount >= 2 {
			return 0
		}
	}

	// Convert the number
	for i < strLen && (str[i] >= 48 && str[i] <= 57) {
		result = result*10 + int(str[i]-'0')
		i++
	}

	return result * sign
}
