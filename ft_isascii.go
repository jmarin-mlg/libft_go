package main

/*
** The isascii() function tests for an ASCII character, which is any character
** between 0 and octal 0177 inclusive (Decimal 0 - 127).
 */
func ftIsASCII(c int) bool {
	return c >= 0 && c <= 127
}
