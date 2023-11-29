package main

/*
** The isdigit() function tests for a decimal digit character.
** Regardless of locale, this includes the following characters only
** The value of the argument must be representable as an unsigned char or the
** value of EOF. The isdigit() and isnumber() functions return zero if the
** character tests false and return non-zero if the character tests true.
 */
func ftIsDigit(c int) bool {
	return c >= '0' && c <= '9'
}
