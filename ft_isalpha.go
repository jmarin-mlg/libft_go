package main

/*
** The isalpha() function tests for any character for which isupper(3) or
** islower(3) is true. The value of the argument must be representable as
** an unsigned char or the value of EOF. The isalpha() function returns zero
** if the character tests false and returns non-zero if the character tests true
 */
func ftIsAlpha(c int) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
