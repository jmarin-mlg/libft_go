package to

import (
	"errors"
	"libft/src/is"
	"libft/src/mem"
	"libft/src/st"
)

/*
** The function converts the initial portion of the string pointed to by str to
** int representation.
** ASCII 9 - 13 and 32 => Spaces ('\t', '\n', '\v', '\f', '\r', ' ')
 */
func FtAtoi(str string) int {
	strLen := st.FtStrlen(str)
	result := 0
	sign := 1
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

	for i < strLen && is.FtIsDigit(int(str[i])) {
		result = (result * 10) + int(str[i]-'0')
		i++
	}

	return result * sign
}

/*
** NumLen calculates the length of the integer representation, considering both
** positive and negative numbers.
**
** Parameters:
**   num: The integer for which the length is calculated.
**
** Returns:
**   The length of the integer representation.
 */
func NumLen(num int) int {
	len := 0

	if num <= 0 {
		len = 1
	}

	for num != 0 {
		len++
		num /= 10
	}

	return len
}

/*
** NumStr converts an unsigned integer to a string and stores it in the provided
** character array starting from the given index.
**
** Parameters:
**   str: The character array to store the string representation.
**   num: The unsigned integer to convert.
**   len: The starting index in the character array to store the string.
**
** Returns:
**   The updated character array with the string representation.
 */
func NumStr(str []byte, num uint, len int) []byte {
	for num > 0 {
		str[len] = byte((num % 10) + '0')
		len--
		num /= 10
	}

	return str
}

/*
** ftItoa converts an integer to a string.
**
** Parameters:
**   n: The integer to convert.
**
** Returns:
**   The string representing the number, or an error if memory allocation fails.
 */
func FtItoa(n int) (string, error) {
	len := NumLen(n)
	res := mem.FtCalloc(uintptr(len+1), 1)

	if res == nil {
		return "", errors.New("Error: memory allocation failed")
	}

	if n == 0 {
		res[0] = '0'
	} else if n < 0 {
		n *= -1
		res[0] = '-'
	}

	return string(NumStr(res, uint(n), len-1)), nil
}

/*
** The tolower() function converts an lower-case letter to the corresponding
** lower-case letter.  The argument must be representable as an unsigned char or
** the value of EOF.
**
** Although the tolower() function uses the current locale, the tolower_l()
** function may be passed a locale directly.
**
** If the argument is an upper-case letter, the tolower() function returns the
** corresponding lower-case letter if there is one; otherwise, the argument is
** returned unchanged.
 */
func FtTolower(c int) int {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}

	return c
}

/*
** The toupper() function converts a lower-case letter to the corresponding
** upper-case letter.  The argument must be representable as an unsigned char or
** the value of EOF.
**
** Although the toupper() function uses the current locale, the toupper_l()
** function may be passed a locale directly.
**
** If the argument is a lower-case letter, the toupper() function returns the
** corresponding upper-case letter if there is one; otherwise, the argument is
** returned unchanged.
 */
func FtToupper(c int) int {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}

	return c
}
