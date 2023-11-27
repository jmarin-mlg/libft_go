package main

/*
** The function computes the length of the string s. The function attempts to
** compute the length of s, but never scans beyond the first maxlen bytes of s.
**
** The function returns the number of characters that precede the terminating
** NUL character.
 */
func ftStrlen(s string) int {
	count := 0

	for range s {
		count++
	}

	return count
}
