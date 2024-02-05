package put

import "os"

/*
** Parameters:
**   c: The character to send.
**   fd: The file descriptor to write to.
 */
func FtPutcharFd(c byte, fd *os.File) {
	fd.Write([]byte{c})
}

/*
** Parameters:
**   s: The string to send.
**   fd: The file descriptor to write to.
 */
func FtPutendlFd(s string, fd *os.File) {
	FtPutstrFd(s, fd)
	FtPutcharFd('\n', fd)
}

/*
** Parameters:
**   n: The number to send.
**   fd: The file descriptor to write to.
 */
func FtPutnbrFd(n int, fd *os.File) {
	if n == -2147483648 {
		FtPutstrFd("-2147483648", fd)
		return
	}

	if n < 0 {
		FtPutcharFd('-', fd)
		n = -n
	}

	divisor := 1

	for n/divisor > 9 {
		divisor *= 10
	}

	for divisor > 0 {
		FtPutcharFd(byte((n/divisor)+'0'), fd)
		n %= divisor
		divisor /= 10
	}
}

/*
** Parameters:
**   s: The string to send.
**   fd: The file descriptor to write to.
 */
func FtPutstrFd(s string, fd *os.File) {
	fd.Write([]byte(s))
}
