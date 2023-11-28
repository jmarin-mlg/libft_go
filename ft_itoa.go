package main

import "errors"

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

func numstr(str []byte, num uint, len int) []byte {
	for num > 0 {
		str[len] = byte((num % 10) + '0')
		len--
		num /= 10
	}
	return str
}

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
