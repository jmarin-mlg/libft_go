package main

/*
** ftStrrchr locates the last occurrence of c in the string s.
**
** It returns the index of the located character, or -1 if the character
** does not appear in the string.
 */
func ftStrrchr(s string, c int) int {
	for i := len(s) - 1; i >= 0; i-- {
		if int(s[i]) == c {
			return i
		}
	}

	if c == 0 {
		return len(s)
	}

	return -1
}
