package main

import "os"

/*
** Parameters:
**   s: The string to send.
**   fd: The file descriptor to write to.
 */
func ftPutendlFd(s string, fd *os.File) {
	ftPutstrFd(s, fd)
	ftPutcharFd('\n', fd)
}
