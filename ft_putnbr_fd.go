package main

import "os"

/*
** Parameters:
**   n: The number to send.
**   fd: The file descriptor to write to.
 */
func ftPutnbrFd(n int, fd *os.File) {
	if n == -2147483648 {
		ftPutstrFd("-2147483648", fd)
		return
	}

	if n < 0 {
		ftPutcharFd('-', fd)
		n = -n
	}

	divisor := 1
	for n/divisor > 9 {
		divisor *= 10
	}

	for divisor > 0 {
		ftPutcharFd(byte((n/divisor)+'0'), fd)
		n %= divisor
		divisor /= 10
	}
}
