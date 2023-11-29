package main

/*
** The isalnum() function tests for any character for which isalpha(3) or
** isdigit(3) is true.  The value of the argument must be representable as an
** unsigned char or the value of EOF. The isalnum() function returns zero if
** the character tests false and returns non-zero if the character tests true.
 */
func ftIsAlnum(c int) bool {
	return ftIsAlpha(c) || ftIsDigit(c)
}
