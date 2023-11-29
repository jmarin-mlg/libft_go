package main

/*
** The toupper() function converts a lower-case letter to the corresponding
** upper-case letter.  The argument must be representable as an unsigned char or
** the value of EOF.
**
** Although the toupper() function uses the current locale, the toupper_l()
** function may be passed a locale directly.
**
** If the argument is a lower-case letter, the toupper() function returns the
** corresponding upper-case letter if there is one; otherwise, the argument is
** returned unchanged.
 */
func ftToupper(c int) int {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}

	return c
}
