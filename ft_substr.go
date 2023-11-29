package main

import "errors"

/*
** ftSubstr returns a substring of s starting from index start with maximum
** length len.
**
** It returns the resulting substring or an error if memory allocation fails.
 */
func ftSubstr(s string, start uint, length int) (string, error) {
	if s == "" {
		return "", nil
	}

	size := ftStrlen(s)

	if start >= uint(size) {
		return "", nil
	}

	if length < 0 {
		length = 0
	}

	if start+uint(length) > uint(size) {
		length = size - int(start)
	}

	res := ftCalloc(uintptr(length+1), 1)
	if res == nil {
		return "", errors.New("Error: memory allocation failed")
	}

	for i := 0; i < length; i++ {
		res[i] = s[start+uint(i)]
	}

	return string(res), nil
}
