package is

/*
** The isalnum() function tests for any character for which isalpha(3) or
** isdigit(3) is true.  The value of the argument must be representable as an
** unsigned char or the value of EOF. The isalnum() function returns zero if
** the character tests false and returns non-zero if the character tests true.
 */
func FtIsAlnum(c int) bool {
	return FtIsAlpha(c) || FtIsDigit(c)
}

/*
** The isalpha() function tests for any character for which isupper(3) or
** islower(3) is true. The value of the argument must be representable as
** an unsigned char or the value of EOF. The isalpha() function returns zero
** if the character tests false and returns non-zero if the character tests true
 */
func FtIsAlpha(c int) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

/*
** The isascii() function tests for an ASCII character, which is any character
** between 0 and octal 0177 inclusive (Decimal 0 - 127).
 */
func FtIsASCII(c int) bool {
	return c >= 0 && c <= 127
}

/*
** The isdigit() function tests for a decimal digit character.
** Regardless of locale, this includes the following characters only
** The value of the argument must be representable as an unsigned char or the
** value of EOF. The isdigit() and isnumber() functions return zero if the
** character tests false and return non-zero if the character tests true.
 */
func FtIsDigit(c int) bool {
	return c >= '0' && c <= '9'
}

/*
** The isprint() function tests for any printing character, including
** space (` ').  The value of the argument must be representable as an unsigned
** char or the value of EOF. The isprint() function returns zero if the
** character tests false and returns non-zero if the character tests true.
 */
func FtIsPrint(c int) bool {
	return c >= 32 && c <= 126
}
