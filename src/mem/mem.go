package mem

import "unsafe"

/*
** The function writes len bytes of value c to the string b.
**
** The function returns its first argument.
 */
func FtMemset(b []byte, c byte, len int) []byte {
	for i := 0; i < len; i++ {
		b[i] = c
	}

	return b
}

/*
** The function writes n zeroed bytes to the string s.
** If n is zero, does nothing.
 */
func FtBzero(s []byte, n int) {
	if n != 0 {
		FtMemset(s, 0, n)
	}
}

// ftCalloc emulates calloc behavior in C
func FtCalloc(count uintptr, size uintptr) []byte {
	sizeRes := count * size
	res := make([]byte, sizeRes)

	if len(res) != int(sizeRes) {
		return nil
	}

	FtBzero(res, int(sizeRes))

	return res
}

/*
** The function locates the first occurrence of c in string s.
**
** The function returns a pointer to the byte located, or nil if no such byte
** exists within n bytes.
 */
func FtMemchr(s []byte, c int, n int) unsafe.Pointer {
	for i := 0; i < n; i++ {
		if s[i] == byte(c) {
			return unsafe.Pointer(&s[i])
		}
	}

	return nil
}

/*
** The function compares byte string s1 against byte string s2. Both strings
** are assumed to be n bytes long.
**
** The function returns zero if the two strings are identical, otherwise returns
** the difference between the first two differing bytes (treated as unsigned
** char values, so that `\200' is greater than `\0', for example). Zero-length
** strings are always identical. This behavior is not required by C and portable
** code should only depend on the sign of the returned value.
 */
func FtMemcmp(s1 []byte, s2 []byte, n int) int {
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			return int(s1[i]) - int(s2[i])
		}
	}

	return 0
}

/*
** The function copies n bytes from memory area src to memory area dst.
** If dst and src overlap, behavior is undefined.
**
** The function returns the original value of dst.
 */
func FtMemcpy(dst []byte, src []byte, n int) []byte {
	if n == 0 || &dst[0] == &src[0] {
		return dst
	}

	for i := 0; i < n; i++ {
		dst[i] = src[i]
	}

	return dst
}

/*
** The function copies len bytes from string src to string dst. The two strings
** may overlap; the copy is always done in a non-destructive manner.
**
** The function returns the original value of dst.
 */
func FtMemmove(dst []byte, src []byte, len int) []byte {
	if len == 0 || &dst[0] == &src[0] {
		return dst
	}

	if uintptr(unsafe.Pointer(&dst[0])) > uintptr(unsafe.Pointer(&src[0])) {
		for i := len - 1; i >= 0; i-- {
			dst[i] = src[i]
		}
	} else {
		FtMemcpy(dst, src, len)
	}

	return dst
}
