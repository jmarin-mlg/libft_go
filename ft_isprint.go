package main

/*
** The isprint() function tests for any printing character, including
** space (` ').  The value of the argument must be representable as an unsigned
** char or the value of EOF. The isprint() function returns zero if the
** character tests false and returns non-zero if the character tests true.
 */
func ftIsPrint(c int) bool {
	return c >= 32 && c <= 126
}
