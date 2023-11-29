package main

import "unsafe"

/*
** The function locates the first occurrence of c in string s.
**
** The function returns a pointer to the byte located, or nil if no such byte
** exists within n bytes.
 */
func ftMemchr(s []byte, c int, n int) unsafe.Pointer {
	for i := 0; i < n; i++ {
		if s[i] == byte(c) {
			return unsafe.Pointer(&s[i])
		}
	}

	return nil
}
