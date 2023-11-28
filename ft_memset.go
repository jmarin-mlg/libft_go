package main

/*
** The function writes len bytes of value c to the string b.
**
** The function returns its first argument.
 */
func ftMemset(b []byte, c byte, len int) []byte {
	for i := 0; i < len; i++ {
		b[i] = c
	}
	return b
}
