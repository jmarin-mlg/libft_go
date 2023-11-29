package main

import "unsafe"

/*
** The function copies len bytes from string src to string dst. The two strings
** may overlap; the copy is always done in a non-destructive manner.
**
** The function returns the original value of dst.
 */
func ftMemmove(dst []byte, src []byte, len int) []byte {
	if len == 0 || &dst[0] == &src[0] {
		return dst
	}

	if uintptr(unsafe.Pointer(&dst[0])) > uintptr(unsafe.Pointer(&src[0])) {
		for i := len - 1; i >= 0; i-- {
			dst[i] = src[i]
		}
	} else {
		ftMemcpy(dst, src, len)
	}

	return dst
}
