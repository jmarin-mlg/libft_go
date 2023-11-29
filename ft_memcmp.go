package main

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
func ftMemcmp(s1 []byte, s2 []byte, n int) int {
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			return int(s1[i]) - int(s2[i])
		}
	}

	return 0
}
