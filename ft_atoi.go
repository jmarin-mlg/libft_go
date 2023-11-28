package main

/*
** The function converts the initial portion of the string pointed to by str to
** int representation.
** ASCII 9 - 13 and 32 => Spaces ('\t', '\n', '\v', '\f', '\r', ' ')
 */
func ftAtoi(str string) int {
	sign := 1
	strLen := ftStrlen(str)
	result := 0
	i := 0

	for i < strLen && ((str[i] >= 9 && str[i] <= 13) || str[i] == 32) {
		i++
	}

	if i < strLen && (str[i] == '+' || str[i] == '-') {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}

	for i < strLen && ftIsDigit(int(str[i])) {
		result = (result * 10) + int(str[i]-'0')
		i++
	}

	return result * sign
}
