package main

/*
** The tolower() function converts an lower-case letter to the corresponding
** lower-case letter.  The argument must be representable as an unsigned char or
** the value of EOF.
**
** Although the tolower() function uses the current locale, the tolower_l()
** function may be passed a locale directly.
**
** If the argument is an upper-case letter, the tolower() function returns the
** corresponding lower-case letter if there is one; otherwise, the argument is
** returned unchanged.
 */
func ftTolower(c int) int {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}

	return c
}
