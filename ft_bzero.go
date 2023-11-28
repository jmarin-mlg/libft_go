package main

/*
** The function writes n zeroed bytes to the string s.
** If n is zero, does nothing.
 */
func ftBzero(s []byte, n int) {
	if n != 0 {
		ftMemset(s, 0, n)
	}
}
