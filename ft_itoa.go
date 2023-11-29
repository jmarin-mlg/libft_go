package main

import "errors"

/*
** numlen calculates the length of the integer representation, considering both
** positive and negative numbers.
**
** Parameters:
**   num: The integer for which the length is calculated.
**
** Returns:
**   The length of the integer representation.
 */
func numlen(num int) int {
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
** numstr converts an unsigned integer to a string and stores it in the provided
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
func numstr(str []byte, num uint, len int) []byte {
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
func ftItoa(n int) (string, error) {
	len := numlen(n)
	res := ftCalloc(uintptr(len+1), 1)

	if res == nil {
		return "", errors.New("Error: memory allocation failed")
	}

	if n == 0 {
		res[0] = '0'
	} else if n < 0 {
		n *= -1
		res[0] = '-'
	}

	return string(numstr(res, uint(n), len-1)), nil
}
