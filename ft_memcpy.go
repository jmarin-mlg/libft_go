package main

/*
** The function copies n bytes from memory area src to memory area dst.
** If dst and src overlap, behavior is undefined.
**
** The function returns the original value of dst.
 */
func ftMemcpy(dst []byte, src []byte, n int) []byte {
	if n == 0 || &dst[0] == &src[0] {
		return dst
	}

	for i := 0; i < n; i++ {
		dst[i] = src[i]
	}

	return dst
}
