package main

/*
** ftStrchr locates the first occurrence of c in the string s.
**
** It returns the index of the located character, or -1 if the character
** does not appear in the string.
 */
func ftStrchr(s string, c int) int {
	for i, char := range s {
		if int(char) == c {
			return i
		}
	}

	if c == 0 {
		return len(s)
	}

	return -1
}
